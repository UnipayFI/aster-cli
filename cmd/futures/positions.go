package futures

import (
	"log"

	"github.com/UnipayFI/aster-cli/printer"
	asterfutures "github.com/UnipayFI/go-aster/v3/futures"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var (
	positionCmd = &cobra.Command{
		Use:   "position",
		Short: "Position management commands",
		Long:  `Manage positions: list, risk, mode, margin, ADL quantile, etc.`,
	}

	positionListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List positions",
		Long: `Get current account's all positions.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#account-information-v3-user_data`,
		Run: listPositions,
	}

	positionRiskCmd = &cobra.Command{
		Use:     "risk",
		Aliases: []string{"r"},
		Short:   "Show position risk",
		Long: `Get current position information (only symbols with positions will be returned).

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#position-information-v3-user_data`,
		Run: showPositionRisk,
	}

	positionModeCmd = &cobra.Command{
		Use:   "mode",
		Short: "Manage position mode (Hedge/One-way)",
		Long:  `Get or change the position mode. Hedge mode allows both LONG and SHORT positions; One-way mode allows only one position direction.`,
	}

	positionModeGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get current position mode",
		Long: `Get current position mode.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#get-current-position-modeuser_data`,
		Run: getPositionMode,
	}

	dualSidePosition   bool
	positionModeSetCmd = &cobra.Command{
		Use:   "set",
		Short: "Set position mode",
		Long: `Change the position mode. Use --dualSidePosition=true for Hedge mode or --dualSidePosition=false for One-way mode.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#change-position-modetrade`,
		Run: setPositionModeFunc,
	}

	positionMarginHistorySymbol    string
	positionMarginHistoryType      int
	positionMarginHistoryStartTime string
	positionMarginHistoryEndTime   string
	positionMarginHistoryLimit     int
	positionMarginHistoryCmd       = &cobra.Command{
		Use:   "margin-history",
		Short: "Query position margin change history",
		Long: `Get the position margin change history for a symbol.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#get-position-margin-change-history-trade`,
		Run: showPositionMarginHistory,
	}

	positionAdlQuantileSymbol string
	positionAdlQuantileCmd    = &cobra.Command{
		Use:   "adl-quantile",
		Short: "Query ADL quantile estimation",
		Long: `Get ADL (Auto-Deleveraging) quantile estimation for positions.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#position-adl-quantile-estimation-user_data`,
		Run: showAdlQuantile,
	}

	positionMarginCmd = &cobra.Command{
		Use:   "set-margin",
		Short: "Modify isolated position margin",
		Long: `Add or reduce isolated position margin.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#modify-isolated-position-margin-trade`,
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
	positionRiskCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol")

	positionModeSetCmd.Flags().BoolVar(&dualSidePosition, "dualSidePosition", false, "true: Hedge Mode; false: One-way Mode")
	positionModeSetCmd.MarkFlagRequired("dualSidePosition")
	positionModeCmd.AddCommand(positionModeGetCmd, positionModeSetCmd)

	positionMarginHistoryCmd.Flags().StringVarP(&positionMarginHistorySymbol, "symbol", "s", "", "Trading pair symbol (required)")
	positionMarginHistoryCmd.Flags().IntVarP(&positionMarginHistoryType, "type", "t", 0, "Margin type: 1 for Add, 2 for Reduce")
	positionMarginHistoryCmd.Flags().StringVarP(&positionMarginHistoryStartTime, "startTime", "a", "", "Start time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	positionMarginHistoryCmd.Flags().StringVarP(&positionMarginHistoryEndTime, "endTime", "e", "", "End time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	positionMarginHistoryCmd.Flags().IntVarP(&positionMarginHistoryLimit, "limit", "l", 500, "Number of results (default 500)")
	positionMarginHistoryCmd.MarkFlagRequired("symbol")

	positionAdlQuantileCmd.Flags().StringVarP(&positionAdlQuantileSymbol, "symbol", "s", "", "Trading pair symbol (optional)")

	positionMarginCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol (required)")
	positionMarginCmd.Flags().StringP("positionSide", "p", "BOTH", "Position side: BOTH, LONG, or SHORT")
	positionMarginCmd.Flags().StringP("amount", "a", "", "Margin amount (decimal string)")
	positionMarginCmd.Flags().StringP("type", "t", "ADD", "Margin type: ADD or REDUCE")
	positionMarginCmd.MarkFlagRequired("symbol")
	positionMarginCmd.MarkFlagRequired("amount")

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
	client := newClient()
	positions, err := client.GetPositions()
	if err != nil {
		log.Fatalf("futures position list error: %v", err)
	}
	printer.Print(&positions)
}

func showPositionRisk(cmd *cobra.Command, _ []string) {
	client := newClient()
	symbol, _ := cmd.Flags().GetString("symbol")
	positions, err := client.GetPositionRisk(symbol)
	if err != nil {
		log.Fatalf("futures position risk error: %v", err)
	}
	printer.Print(&positions)
}

func getPositionMode(cmd *cobra.Command, args []string) {
	client := newClient()
	dualSide, err := client.GetPositionMode()
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(map[string]bool{"dualSidePosition": dualSide})
}

func setPositionModeFunc(cmd *cobra.Command, args []string) {
	if !cmd.Flags().Changed("dualSidePosition") {
		log.Fatal("Please specify --dualSidePosition")
	}

	client := newClient()
	resp, err := client.ChangePositionMode(dualSidePosition)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(resp)
}

func showPositionMarginHistory(cmd *cobra.Command, args []string) {
	client := newClient()
	startTime, err := parseTimeFlag("--startTime", positionMarginHistoryStartTime)
	if err != nil {
		log.Fatal(err)
	}
	endTime, err := parseTimeFlag("--endTime", positionMarginHistoryEndTime)
	if err != nil {
		log.Fatal(err)
	}
	history, err := client.GetPositionMarginHistory(positionMarginHistorySymbol, asterfutures.PositionMarginType(positionMarginHistoryType), startTime, endTime, positionMarginHistoryLimit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(history)
}

func showAdlQuantile(cmd *cobra.Command, args []string) {
	client := newClient()
	quantiles, err := client.GetAdlQuantile(positionAdlQuantileSymbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(quantiles)
}

func setPositionMargin(cmd *cobra.Command, _ []string) {
	client := newClient()
	symbol, _ := cmd.Flags().GetString("symbol")
	positionSide, _ := cmd.Flags().GetString("positionSide")
	amountRaw, _ := cmd.Flags().GetString("amount")
	typ, _ := cmd.Flags().GetString("type")

	amount, err := decimal.NewFromString(amountRaw)
	if err != nil {
		log.Fatalf("invalid amount: %v", err)
	}

	action := asterfutures.PositionMarginAdd
	if typ == "REDUCE" {
		action = asterfutures.PositionMarginReduce
	}
	resp, err := client.ModifyPositionMargin(symbol, positionSide, amount, action)
	if err != nil {
		log.Fatalf("futures position margin set error: %v", err)
	}
	printer.Print(resp)
}
