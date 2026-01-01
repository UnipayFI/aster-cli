package futures

import (
	"fmt"
	"log"
	"time"

	"github.com/UnipayFI/aster-cli/common"
	"github.com/UnipayFI/aster-cli/config"
	"github.com/UnipayFI/aster-cli/exchange"
	"github.com/UnipayFI/aster-cli/exchange/futures"
	"github.com/UnipayFI/aster-cli/printer"
	"github.com/spf13/cobra"
)

var (
	positionCmd = &cobra.Command{
		Use:   "position",
		Short: "Position management commands",
		Long:  `Manage positions: list, risk, mode, margin, ADL quantile, etc.`,
	}

	// position list
	positionListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List positions",
		Long:    `Get current account's all positions.`,
		Run:     listPositions,
	}

	// position risk
	positionRiskCmd = &cobra.Command{
		Use:     "risk",
		Aliases: []string{"r"},
		Short:   "Show position risk",
		Long:    `Get current position information (only symbols with positions will be returned).`,
		Run:     showPositionRisk,
	}

	// position mode
	positionModeCmd = &cobra.Command{
		Use:   "mode",
		Short: "Manage position mode (Hedge/One-way)",
		Long:  `Get or change the position mode. Hedge mode allows both LONG and SHORT positions; One-way mode allows only one position direction.`,
	}

	positionModeGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get current position mode",
		Run:   getPositionMode,
	}

	dualSidePosition   bool
	positionModeSetCmd = &cobra.Command{
		Use:   "set",
		Short: "Set position mode",
		Long:  `Change the position mode. Use --dualSidePosition=true for Hedge mode or --dualSidePosition=false for One-way mode.`,
		Run:   setPositionModeFunc,
	}

	// position margin-history
	positionMarginHistorySymbol    string
	positionMarginHistoryType      int
	positionMarginHistoryStartTime string
	positionMarginHistoryEndTime   string
	positionMarginHistoryLimit     int
	positionMarginHistoryCmd       = &cobra.Command{
		Use:   "margin-history",
		Short: "Query position margin change history",
		Long:  `Get the position margin change history for a symbol.`,
		Run:   showPositionMarginHistory,
	}

	// position adl-quantile
	positionAdlQuantileSymbol string
	positionAdlQuantileCmd    = &cobra.Command{
		Use:   "adl-quantile",
		Short: "Query ADL quantile estimation",
		Long:  `Get ADL (Auto-Deleveraging) quantile estimation for positions.`,
		Run:   showAdlQuantile,
	}

	// position set-margin
	positionMarginCmd = &cobra.Command{
		Use:   "set-margin",
		Short: "Modify isolated position margin",
		Long:  `Add or reduce isolated position margin.`,
		PreRun: func(cmd *cobra.Command, args []string) {
			typ, _ := cmd.Flags().GetString("type")
			if typ != "ADD" && typ != "REDUCE" {
				log.Fatalf("type must be ADD or REDUCE")
			}
		},
		Run: setPositionMargin,
	}
)

func InitPositionsCmds() []*cobra.Command {
	// position risk flags
	positionRiskCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol")

	// position mode flags
	positionModeSetCmd.Flags().BoolVar(&dualSidePosition, "dualSidePosition", false, "true: Hedge Mode; false: One-way Mode")
	positionModeSetCmd.MarkFlagRequired("dualSidePosition")
	positionModeCmd.AddCommand(positionModeGetCmd, positionModeSetCmd)

	// position margin-history flags
	positionMarginHistoryCmd.Flags().StringVarP(&positionMarginHistorySymbol, "symbol", "s", "", "Trading pair symbol (required)")
	positionMarginHistoryCmd.Flags().IntVarP(&positionMarginHistoryType, "type", "t", 0, "Margin type: 1 for Add, 2 for Reduce")
	positionMarginHistoryCmd.Flags().StringVarP(&positionMarginHistoryStartTime, "startTime", "a", "", "Start time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	positionMarginHistoryCmd.Flags().StringVarP(&positionMarginHistoryEndTime, "endTime", "e", "", "End time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	positionMarginHistoryCmd.Flags().IntVarP(&positionMarginHistoryLimit, "limit", "l", 500, "Number of results (default 500)")
	positionMarginHistoryCmd.MarkFlagRequired("symbol")

	// position adl-quantile flags
	positionAdlQuantileCmd.Flags().StringVarP(&positionAdlQuantileSymbol, "symbol", "s", "", "Trading pair symbol (optional)")

	// position set-margin flags
	positionMarginCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol (required)")
	positionMarginCmd.Flags().StringP("positionSide", "p", "BOTH", "Position side: BOTH, LONG, or SHORT")
	positionMarginCmd.Flags().Float64P("amount", "a", 0, "Margin amount")
	positionMarginCmd.Flags().StringP("type", "t", "ADD", "Margin type: ADD or REDUCE")
	positionMarginCmd.MarkFlagRequired("symbol")
	positionMarginCmd.MarkFlagRequired("amount")

	// Add all subcommands to position
	positionCmd.AddCommand(
		positionListCmd,
		positionRiskCmd,
		positionModeCmd,
		positionMarginHistoryCmd,
		positionAdlQuantileCmd,
		positionMarginCmd,
	)

	return []*cobra.Command{positionCmd}
}

func listPositions(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	positions, err := client.GetPositions()
	if err != nil {
		log.Fatalf("futures position list error: %v", err)
	}
	printer.PrintTable(&positions)
}

func showPositionRisk(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	positions, err := client.GetPositionRisk(symbol)
	if err != nil {
		log.Fatalf("futures position risk error: %v", err)
	}
	printer.PrintTable(&positions)
}

func getPositionMode(cmd *cobra.Command, args []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	dualSide, err := client.GetPositionMode()
	if err != nil {
		log.Fatal(err)
	}
	if dualSide {
		fmt.Println("Position Mode: Hedge Mode (Dual Side Position)")
	} else {
		fmt.Println("Position Mode: One-way Mode (Single Side Position)")
	}
}

func setPositionModeFunc(cmd *cobra.Command, args []string) {
	if !cmd.Flags().Changed("dualSidePosition") {
		log.Fatal("Please specify --dualSidePosition")
	}

	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	err := client.ChangePositionMode(dualSidePosition)
	if err != nil {
		log.Fatal(err)
	}

	if dualSidePosition {
		fmt.Println("Position mode changed to: Hedge Mode (Dual Side Position)")
	} else {
		fmt.Println("Position mode changed to: One-way Mode (Single Side Position)")
	}
}

func showPositionMarginHistory(cmd *cobra.Command, args []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	var startTime, endTime time.Time
	if t, ok, err := common.ParseTimeFlag("--startTime", positionMarginHistoryStartTime); err != nil {
		log.Fatal(err)
	} else if ok {
		startTime = t
	}
	if t, ok, err := common.ParseTimeFlag("--endTime", positionMarginHistoryEndTime); err != nil {
		log.Fatal(err)
	} else if ok {
		endTime = t
	}
	history, err := client.GetPositionMarginHistory(positionMarginHistorySymbol, positionMarginHistoryType, startTime, endTime, positionMarginHistoryLimit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(history)
}

func showAdlQuantile(cmd *cobra.Command, args []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	quantiles, err := client.GetAdlQuantile(positionAdlQuantileSymbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(quantiles)
}

func setPositionMargin(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	positionSide, _ := cmd.Flags().GetString("positionSide")
	amount, _ := cmd.Flags().GetFloat64("amount")
	typ, _ := cmd.Flags().GetString("type")
	var t int
	if typ == "ADD" {
		t = 1
	} else {
		t = 2
	}
	err := client.ModifyPositionMargin(symbol, positionSide, amount, t)
	if err != nil {
		log.Fatalf("futures position margin set error: %v", err)
	}
	fmt.Printf("%s %s position %s %.6f\n", symbol, positionSide, typ, amount)
}
