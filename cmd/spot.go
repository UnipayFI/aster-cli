package cmd

import (
	"github.com/UnipayFI/aster-cli/cmd/spot"
	"github.com/spf13/cobra"
)

var (
	SpotCmd = &cobra.Command{
		Use:   "spot",
		Short: "Spot trading commands",
	}
)

func init() {
	SpotCmd.AddCommand(spot.InitAccountCmds()...)
	SpotCmd.AddCommand(spot.InitOrderCmds()...)
	SpotCmd.AddCommand(spot.InitTradeCmds()...)
	SpotCmd.AddCommand(spot.InitCommissionRateCmds()...)
	RootCmd.AddCommand(SpotCmd)
}
