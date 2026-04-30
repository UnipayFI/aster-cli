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
		Long: `Get current account information.

Docs Link: https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#account-information-user_data`,
		Run: showAccount,
	}

	balanceCmd = &cobra.Command{
		Use:   "balance",
		Short: "Show account balances",
		Long: `Get current account balances (non-zero only).

Docs Link: https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#account-information-user_data`,
		Run: showBalance,
	}
)

func InitAccountCmds() []*cobra.Command {
	return []*cobra.Command{accountCmd, balanceCmd}
}

func newClient() spot.Client {
	return spot.Client{Client: exchange.NewClient(config.Config.APIAddress, config.Config.APIPrivateKey, config.Config.ChainID)}
}

func showAccount(cmd *cobra.Command, args []string) {
	client := newClient()
	account, err := client.GetAccountInfo()
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(account)
}

func showBalance(cmd *cobra.Command, args []string) {
	client := newClient()
	account, err := client.GetAccountInfo()
	if err != nil {
		log.Fatal(err)
	}
	balances := spot.FilterNonZeroBalances(account.Balances)
	printer.Print(balances)
}
