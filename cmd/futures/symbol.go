package futures

import (
	"fmt"
	"log"

	"github.com/UnipayFI/aster-cli/config"
	"github.com/UnipayFI/aster-cli/exchange"
	"github.com/UnipayFI/aster-cli/exchange/futures"
	"github.com/UnipayFI/aster-cli/printer"
	asterfutures "github.com/UnipayFI/go-aster/futures"
	"github.com/spf13/cobra"
)

var (
	symbolCmd = &cobra.Command{
		Use:   "symbol",
		Short: "Symbol configuration commands",
		Long:  `Manage symbol configuration: leverage, margin type, leverage brackets.`,
	}

	// symbol set-leverage
	symbolSetLeverageCmd = &cobra.Command{
		Use:     "set-leverage",
		Aliases: []string{"leverage"},
		Short:   "Change initial leverage",
		Long:    `Change initial leverage for a symbol.`,
		Run:     setSymbolLeverage,
	}

	// symbol set-margin-type
	symbolSetMarginTypeCmd = &cobra.Command{
		Use:     "set-margin-type",
		Aliases: []string{"margin-type"},
		Short:   "Set margin type",
		Long:    `Change symbol level margin type (ISOLATED or CROSSED).`,
		Run:     setSymbolMarginType,
	}

	// symbol leverage-bracket
	symbolLeverageBracketSymbol string
	symbolLeverageBracketCmd    = &cobra.Command{
		Use:   "leverage-bracket",
		Short: "Query leverage bracket information",
		Long:  `Get notional and leverage bracket information for a symbol or all symbols.`,
		Run:   showSymbolLeverageBracket,
	}
)

func InitSymbolCmds() []*cobra.Command {
	// symbol set-leverage flags
	symbolSetLeverageCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol (required)")
	symbolSetLeverageCmd.Flags().IntP("leverage", "l", 0, "Leverage value (required)")
	symbolSetLeverageCmd.MarkFlagRequired("symbol")
	symbolSetLeverageCmd.MarkFlagRequired("leverage")

	// symbol set-margin-type flags
	symbolSetMarginTypeCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol (required)")
	symbolSetMarginTypeCmd.Flags().StringP("marginType", "m", "", "Margin type: ISOLATED or CROSSED (required)")
	symbolSetMarginTypeCmd.MarkFlagRequired("symbol")
	symbolSetMarginTypeCmd.MarkFlagRequired("marginType")

	// symbol leverage-bracket flags
	symbolLeverageBracketCmd.Flags().StringVarP(&symbolLeverageBracketSymbol, "symbol", "s", "", "Trading pair symbol (optional)")

	symbolCmd.AddCommand(
		symbolSetLeverageCmd,
		symbolSetMarginTypeCmd,
		symbolLeverageBracketCmd,
	)
	return []*cobra.Command{symbolCmd}
}

func setSymbolLeverage(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	leverage, _ := cmd.Flags().GetInt("leverage")
	resp, err := client.SetLeverage(symbol, leverage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("leverage set to %d for %s (max notional: %s)\n", resp.Leverage, resp.Symbol, resp.MaxNotionalValue)
}

func setSymbolMarginType(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	marginType, _ := cmd.Flags().GetString("marginType")
	err := client.SetMarginType(symbol, asterfutures.MarginType(marginType))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("margin type set to", marginType)
}

func showSymbolLeverageBracket(cmd *cobra.Command, args []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	brackets, err := client.GetLeverageBrackets(symbolLeverageBracketSymbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(brackets)
}
