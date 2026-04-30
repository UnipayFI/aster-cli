package spot

import (
	"log"

	"github.com/UnipayFI/aster-cli/printer"
	"github.com/spf13/cobra"
)

var (
	commissionRateSymbol string

	commissionRateCmd = &cobra.Command{
		Use:   "commission-rate",
		Short: "Get commission rate for a symbol",
		Long: `Query the commission rate for a specific trading pair.

Docs Link: https://asterdex.github.io/aster-api-website/spot-v3/market-data/#get-symbol-fees`,
		Run: showCommissionRate,
	}
)

func InitCommissionRateCmds() []*cobra.Command {
	commissionRateCmd.Flags().StringVarP(&commissionRateSymbol, "symbol", "s", "", "Trading pair symbol (required)")
	commissionRateCmd.MarkFlagRequired("symbol")
	return []*cobra.Command{commissionRateCmd}
}

func showCommissionRate(cmd *cobra.Command, args []string) {
	client := newClient()
	rate, err := client.GetCommissionRate(commissionRateSymbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(rate)
}
