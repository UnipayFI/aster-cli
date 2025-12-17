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
	tradeCmd = &cobra.Command{
		Use:   "trade",
		Short: "Query user trades",
	}

	tradeListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List user trades",
		Long:    `Get trades for a specific account and symbol.`,
		Run:     tradeList,
	}
)

func InitTradeCmds() []*cobra.Command {
	tradeListCmd.Flags().StringP("symbol", "s", "", "symbol (required)")
	tradeListCmd.Flags().Int64P("orderId", "o", 0, "order ID to filter trades")
	tradeListCmd.Flags().Int64P("fromId", "f", 0, "trade ID to fetch from")
	tradeListCmd.Flags().IntP("limit", "l", 500, "limit, max 1000")
	tradeListCmd.MarkFlagRequired("symbol")

	tradeCmd.AddCommand(tradeListCmd)
	return []*cobra.Command{tradeCmd}
}

func tradeList(cmd *cobra.Command, args []string) {
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orderId, _ := cmd.Flags().GetInt64("orderId")
	fromId, _ := cmd.Flags().GetInt64("fromId")
	limit, _ := cmd.Flags().GetInt("limit")

	trades, err := client.GetUserTrades(symbol, orderId, fromId, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(trades)
}
