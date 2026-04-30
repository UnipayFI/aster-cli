package wallet

import (
	"log"

	"github.com/UnipayFI/aster-cli/config"
	"github.com/UnipayFI/aster-cli/exchange"
	"github.com/UnipayFI/aster-cli/exchange/wallet"
	"github.com/UnipayFI/aster-cli/printer"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var (
	transferCmd = &cobra.Command{
		Use:   "transfer",
		Short: "Transfer assets between spot and futures",
		PreRun: func(cmd *cobra.Command, args []string) {
			kindType, _ := cmd.Flags().GetString("kindType")
			asset, _ := cmd.Flags().GetString("asset")
			amount, _ := cmd.Flags().GetString("amount")
			if kindType == "" || asset == "" {
				log.Fatal("kindType, asset are required")
			}
			amt, err := decimal.NewFromString(amount)
			if err != nil {
				log.Fatalf("invalid amount: %v", err)
			}
			if !amt.IsPositive() {
				log.Fatal("amount must be greater than 0")
			}
			if kindType != "SPOT_FUTURE" && kindType != "FUTURE_SPOT" {
				log.Fatal("kindType must be SPOT_FUTURE or FUTURE_SPOT")
			}
		},
		Long: `Transfer assets between spot and futures wallet.

Supported transfer types:
  - SPOT_FUTURE: Transfer from spot wallet to futures wallet
  - FUTURE_SPOT: Transfer from futures wallet to spot wallet

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#transfer-between-futures-and-spot-transfer`,
		Run: doTransfer,
	}
)

func InitTransferCmds() []*cobra.Command {
	transferCmd.Flags().StringP("kindType", "t", "", "kindType: SPOT_FUTURE or FUTURE_SPOT")
	transferCmd.Flags().StringP("asset", "a", "", "asset to transfer (e.g., USDT)")
	transferCmd.Flags().StringP("amount", "m", "", "amount to transfer (decimal string)")
	return []*cobra.Command{transferCmd}
}

func doTransfer(cmd *cobra.Command, args []string) {
	kindType, _ := cmd.Flags().GetString("kindType")
	asset, _ := cmd.Flags().GetString("asset")
	amountRaw, _ := cmd.Flags().GetString("amount")
	amount, err := decimal.NewFromString(amountRaw)
	if err != nil {
		log.Fatalf("invalid amount: %v", err)
	}

	client := wallet.Client{Client: exchange.NewClient(config.Config.APIAddress, config.Config.APIPrivateKey, config.Config.ChainID)}
	result, err := client.Transfer(kindType, asset, amount)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(&wallet.TransferResult{TransferResponse: result})
}
