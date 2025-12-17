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
	accountCmd = &cobra.Command{
		Use:   "account",
		Short: "Show account info",
		Long:  `Get current account information.`,
		Run:   showAccount,
	}

	balanceCmd = &cobra.Command{
		Use:   "balance",
		Short: "Show account balances",
		Long:  `Get current account balances (non-zero only).`,
		Run:   showBalance,
	}
)

func InitAccountCmds() []*cobra.Command {
	return []*cobra.Command{accountCmd, balanceCmd}
}

func showAccount(cmd *cobra.Command, args []string) {
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	account, err := client.GetAccountInfo()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(account)
}

func showBalance(cmd *cobra.Command, args []string) {
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	account, err := client.GetAccountInfo()
	if err != nil {
		log.Fatal(err)
	}
	balances := spot.FilterNonZeroBalances(account.Balances)
	printer.PrintTable(balances)
}
