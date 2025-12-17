package wallet

import (
	"fmt"
	"log"

	"github.com/UnipayFI/aster-cli/config"
	"github.com/UnipayFI/aster-cli/exchange"
	"github.com/UnipayFI/aster-cli/exchange/wallet"
	"github.com/UnipayFI/aster-cli/printer"
	"github.com/spf13/cobra"
)

var (
	transferCmd = &cobra.Command{
		Use:   "transfer",
		Short: "Transfer assets between spot and futures",
		PreRun: func(cmd *cobra.Command, args []string) {
			transferType, _ := cmd.Flags().GetString("type")
			asset, _ := cmd.Flags().GetString("asset")
			amount, _ := cmd.Flags().GetFloat64("amount")
			if transferType == "" || asset == "" {
				log.Fatal("type, asset are required")
			}
			if amount <= 0 {
				log.Fatal("amount must be greater than 0")
			}
			if transferType != "SPOT_FUTURE" && transferType != "FUTURE_SPOT" {
				log.Fatal("type must be SPOT_FUTURE or FUTURE_SPOT")
			}
		},
		Long: `Transfer assets between spot and futures wallet.

Supported transfer types:
  - SPOT_FUTURE: Transfer from spot wallet to futures wallet
  - FUTURE_SPOT: Transfer from futures wallet to spot wallet`,
		Run: doTransfer,
	}
)

func InitTransferCmds() []*cobra.Command {
	transferCmd.Flags().StringP("type", "t", "", "transfer type: SPOT_FUTURE or FUTURE_SPOT")
	transferCmd.Flags().StringP("asset", "a", "", "asset to transfer (e.g., USDT)")
	transferCmd.Flags().Float64P("amount", "m", 0, "amount to transfer")
	return []*cobra.Command{transferCmd}
}

func doTransfer(cmd *cobra.Command, args []string) {
	transferType, _ := cmd.Flags().GetString("type")
	asset, _ := cmd.Flags().GetString("asset")
	amount, _ := cmd.Flags().GetFloat64("amount")

	client := wallet.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	result, err := client.Transfer(transferType, asset, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transfer successful!\n")
	printer.PrintTable(&wallet.TransferResult{WalletTransferResponse: result})
}
