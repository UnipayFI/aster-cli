package cmd

import (
	"github.com/UnipayFI/aster-cli/cmd/wallet"
	"github.com/spf13/cobra"
)

var (
	walletCmd = &cobra.Command{
		Use:   "wallet",
		Short: "Wallet commands",
	}
)

func init() {
	walletCmd.AddCommand(wallet.InitTransferCmds()...)
	RootCmd.AddCommand(walletCmd)
}
