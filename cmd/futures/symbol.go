package futures

import (
	"log"

	"github.com/UnipayFI/aster-cli/printer"
	asterfutures "github.com/UnipayFI/go-aster/v3/futures"
	"github.com/spf13/cobra"
)

var (
	symbolCmd = &cobra.Command{
		Use:   "symbol",
		Short: "Symbol configuration commands",
		Long:  `Manage symbol configuration: leverage, margin type, leverage brackets.`,
	}

	symbolSetLeverageCmd = &cobra.Command{
		Use:     "set-leverage",
		Aliases: []string{"leverage"},
		Short:   "Change initial leverage",
		Long: `Change initial leverage for a symbol.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#change-initial-leverage-trade`,
		Run: setSymbolLeverage,
	}

	symbolSetMarginTypeCmd = &cobra.Command{
		Use:     "set-margin-type",
		Aliases: []string{"margin-type"},
		Short:   "Set margin type",
		Long: `Change symbol level margin type (ISOLATED or CROSSED).

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#change-margin-type-trade`,
		Run: setSymbolMarginType,
	}

	symbolLeverageBracketSymbol string
	symbolLeverageBracketCmd    = &cobra.Command{
		Use:   "leverage-bracket",
		Short: "Query leverage bracket information",
		Long: `Get notional and leverage bracket information for a symbol or all symbols.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#notional-and-leverage-brackets-user_data`,
		Run: showSymbolLeverageBracket,
	}
)

func InitSymbolCmds() []*cobra.Command {
	symbolSetLeverageCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol (required)")
	symbolSetLeverageCmd.Flags().IntP("leverage", "l", 0, "Leverage value (required)")
	symbolSetLeverageCmd.MarkFlagRequired("symbol")
	symbolSetLeverageCmd.MarkFlagRequired("leverage")

	symbolSetMarginTypeCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol (required)")
	symbolSetMarginTypeCmd.Flags().StringP("marginType", "m", "", "Margin type: ISOLATED or CROSSED (required)")
	symbolSetMarginTypeCmd.MarkFlagRequired("symbol")
	symbolSetMarginTypeCmd.MarkFlagRequired("marginType")

	symbolLeverageBracketCmd.Flags().StringVarP(&symbolLeverageBracketSymbol, "symbol", "s", "", "Trading pair symbol (optional)")

	symbolCmd.AddCommand(
		symbolSetLeverageCmd,
		symbolSetMarginTypeCmd,
		symbolLeverageBracketCmd,
	)
	return []*cobra.Command{symbolCmd}
}

func setSymbolLeverage(cmd *cobra.Command, _ []string) {
	client := newClient()
	symbol, _ := cmd.Flags().GetString("symbol")
	leverage, _ := cmd.Flags().GetInt("leverage")
	resp, err := client.SetLeverage(symbol, leverage)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(resp)
}

func setSymbolMarginType(cmd *cobra.Command, _ []string) {
	client := newClient()
	symbol, _ := cmd.Flags().GetString("symbol")
	marginType, _ := cmd.Flags().GetString("marginType")
	resp, err := client.SetMarginType(symbol, asterfutures.MarginType(marginType))
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(resp)
}

func showSymbolLeverageBracket(cmd *cobra.Command, args []string) {
	client := newClient()
	brackets, err := client.GetLeverageBrackets(symbolLeverageBracketSymbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(brackets)
}
