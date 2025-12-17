package spot

import (
	"log"

	"github.com/UnipayFI/aster-cli/config"
	"github.com/UnipayFI/aster-cli/exchange"
	"github.com/UnipayFI/aster-cli/exchange/spot"
	"github.com/UnipayFI/aster-cli/printer"
	"github.com/spf13/cobra"
)

var (
	commissionRateSymbol string

	commissionRateCmd = &cobra.Command{
		Use:   "commission-rate",
		Short: "Get commission rate for a symbol",
		Long:  `Query the commission rate for a specific trading pair.`,
		Run:   showCommissionRate,
	}
)

func InitCommissionRateCmds() []*cobra.Command {
	commissionRateCmd.Flags().StringVarP(&commissionRateSymbol, "symbol", "s", "", "Trading pair symbol (required)")
	commissionRateCmd.MarkFlagRequired("symbol")
	return []*cobra.Command{commissionRateCmd}
}

func showCommissionRate(cmd *cobra.Command, args []string) {
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	rate, err := client.GetCommissionRate(commissionRateSymbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(rate)
}
