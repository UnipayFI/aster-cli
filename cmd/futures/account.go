package futures

import (
	"log"
	"time"

	"github.com/UnipayFI/aster-cli/common"
	"github.com/UnipayFI/aster-cli/config"
	"github.com/UnipayFI/aster-cli/exchange"
	"github.com/UnipayFI/aster-cli/exchange/futures"
	"github.com/UnipayFI/aster-cli/printer"
	"github.com/spf13/cobra"
)

var (
	accountCmd = &cobra.Command{
		Use:   "account",
		Short: "Account management commands",
		Long:  `Manage account: balances, info, commission rate, income, multi-assets mode.`,
	}

	balancesCmd = &cobra.Command{
		Use:     "balances",
		Aliases: []string{"balance", "b"},
		Short:   "Show account balances",
		Long: `Get current account's balances.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#futures-account-balance-v3-user_data`,
		Run: balances,
	}

	accountInfoCmd = &cobra.Command{
		Use:     "info",
		Aliases: []string{"i"},
		Short:   "Show account info",
		Long: `Query account information.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#account-information-v3-user_data`,
		Run: accountInfo,
	}

	accountCommissionRateCmd = &cobra.Command{
		Use:     "commission-rate",
		Aliases: []string{"cr"},
		Short:   "Show commission rate",
		Long: `Get user commission rate for a symbol.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#user-commission-rate-user_data`,
		Run: showAccountCommissionRate,
	}

	accountIncomeCmd = &cobra.Command{
		Use:   "income",
		Short: "Query income history",
		Long: `Query income history.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#get-income-historyuser_data`,
		Run: showAccountIncome,
	}

	accountMultiAssetsModeCmd = &cobra.Command{
		Use:   "multi-assets-mode",
		Short: "Manage multi-assets mode",
		Long:  `Get or set multi-assets mode.`,
	}

	accountMultiAssetsModeShowCmd = &cobra.Command{
		Use:   "show",
		Short: "Show multi-assets mode",
		Long: `Show current multi-assets margin mode.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#get-current-multi-assets-mode-user_data`,
		Run: showAccountMultiAssetsMode,
	}

	multiAssetsMargin            bool
	accountMultiAssetsModeSetCmd = &cobra.Command{
		Use:   "set",
		Short: "Set multi-assets mode",
		Long: `Change multi-assets mode. Use --multiAssetsMargin=true for Multi-Assets Mode or --multiAssetsMargin=false for Single-Asset Mode.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#change-multi-assets-mode-trade`,
		Run: setAccountMultiAssetsMode,
	}
)

func InitAccountCmds() []*cobra.Command {
	accountCommissionRateCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol (required)")
	accountCommissionRateCmd.MarkFlagRequired("symbol")

	accountIncomeCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol")
	accountIncomeCmd.Flags().StringP("incomeType", "t", "", "Income type")
	accountIncomeCmd.Flags().StringP("startTime", "a", "", "Start time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	accountIncomeCmd.Flags().StringP("endTime", "e", "", "End time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	accountIncomeCmd.Flags().IntP("limit", "l", 100, "Number of results (default 100, max 1000)")

	accountMultiAssetsModeSetCmd.Flags().BoolVar(&multiAssetsMargin, "multiAssetsMargin", false, "true: Multi-Assets Mode; false: Single-Asset Mode")
	accountMultiAssetsModeSetCmd.MarkFlagRequired("multiAssetsMargin")
	accountMultiAssetsModeCmd.AddCommand(accountMultiAssetsModeShowCmd, accountMultiAssetsModeSetCmd)

	accountCmd.AddCommand(
		balancesCmd,
		accountInfoCmd,
		accountCommissionRateCmd,
		accountIncomeCmd,
		accountMultiAssetsModeCmd,
	)
	return []*cobra.Command{accountCmd}
}

func newClient() futures.Client {
	return futures.Client{Client: exchange.NewClient(config.Config.APIAddress, config.Config.APIPrivateKey, config.Config.ChainID)}
}

func balances(cmd *cobra.Command, args []string) {
	client := newClient()
	balances, err := client.GetBalances()
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(balances)
}

func accountInfo(cmd *cobra.Command, args []string) {
	client := newClient()
	account, err := client.GetAccount()
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(&futures.AccountInfo{AccountInfo: account})
}

func showAccountCommissionRate(cmd *cobra.Command, _ []string) {
	client := newClient()
	symbol, _ := cmd.Flags().GetString("symbol")
	commissionRate, err := client.GetCommissionRate(symbol)
	if err != nil {
		log.Fatalf("futures commission rate error: %v", err)
	}
	printer.Print(&commissionRate)
}

func showAccountIncome(cmd *cobra.Command, _ []string) {
	symbol, _ := cmd.Flags().GetString("symbol")
	incomeType, _ := cmd.Flags().GetString("incomeType")
	startTimeRaw, _ := cmd.Flags().GetString("startTime")
	endTimeRaw, _ := cmd.Flags().GetString("endTime")
	limit, _ := cmd.Flags().GetInt("limit")

	startTime, err := parseTimeFlag("--startTime", startTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	endTime, err := parseTimeFlag("--endTime", endTimeRaw)
	if err != nil {
		log.Fatal(err)
	}

	client := newClient()
	income, err := client.GetIncome(symbol, incomeType, startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(&income)
}

func showAccountMultiAssetsMode(cmd *cobra.Command, _ []string) {
	client := newClient()
	multiAssetsMode, err := client.GetMultiAssetsMode()
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(map[string]bool{"multiAssetsMargin": multiAssetsMode})
}

func setAccountMultiAssetsMode(cmd *cobra.Command, _ []string) {
	if !cmd.Flags().Changed("multiAssetsMargin") {
		log.Fatal("Please specify --multiAssetsMargin")
	}

	client := newClient()
	resp, err := client.SetMultiAssetsMode(multiAssetsMargin)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(resp)
}

// parseTimeFlag is a thin wrapper that returns a zero time when the flag was
// not supplied; the underlying error/parsed-state distinction matters only
// for validating user input.
func parseTimeFlag(name, value string) (time.Time, error) {
	t, _, err := common.ParseTimeFlag(name, value)
	return t, err
}
