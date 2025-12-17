package cmd

import (
	"github.com/UnipayFI/aster-cli/cmd/futures"
	"github.com/spf13/cobra"
)

var futuresCmd = &cobra.Command{
	Use:   "futures",
	Short: "Futures trading commands",
	Long:  `Futures trading commands: account, position, order, symbol, funding.`,
}

func init() {
	futuresCmd.AddCommand(futures.InitAccountCmds()...)
	futuresCmd.AddCommand(futures.InitPositionsCmds()...)
	futuresCmd.AddCommand(futures.InitOrderCmds()...)
	futuresCmd.AddCommand(futures.InitSymbolCmds()...)
	futuresCmd.AddCommand(futures.InitFundingCmds()...)
	RootCmd.AddCommand(futuresCmd)
}
