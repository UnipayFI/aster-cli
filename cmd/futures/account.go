package futures

import (
	"fmt"
	"log"

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

	// account balances
	balancesCmd = &cobra.Command{
		Use:     "balances",
		Aliases: []string{"balance", "b"},
		Short:   "Show account balances",
		Long:    `Get current account's balances.`,
		Run:     balances,
	}

	// account info
	accountInfoCmd = &cobra.Command{
		Use:     "info",
		Aliases: []string{"i"},
		Short:   "Show account info",
		Long:    `Query account information.`,
		Run:     accountInfo,
	}

	// account commission-rate
	accountCommissionRateCmd = &cobra.Command{
		Use:     "commission-rate",
		Aliases: []string{"cr"},
		Short:   "Show commission rate",
		Long:    `Get user commission rate for a symbol.`,
		Run:     showAccountCommissionRate,
	}

	// account income
	accountIncomeCmd = &cobra.Command{
		Use:   "income",
		Short: "Query income history",
		Long:  `Query income history.`,
		Run:   showAccountIncome,
	}

	// account multi-assets-mode
	accountMultiAssetsModeCmd = &cobra.Command{
		Use:   "multi-assets-mode",
		Short: "Manage multi-assets mode",
		Long:  `Get or set multi-assets mode.`,
	}

	accountMultiAssetsModeShowCmd = &cobra.Command{
		Use:   "show",
		Short: "Show multi-assets mode",
		Run:   showAccountMultiAssetsMode,
	}

	multiAssetsModeEnable        bool
	multiAssetsModeDisable       bool
	accountMultiAssetsModeSetCmd = &cobra.Command{
		Use:   "set",
		Short: "Set multi-assets mode",
		Long:  `Change multi-assets mode. Use --enable to enable or --disable to disable.`,
		Run:   setAccountMultiAssetsMode,
	}
)

func InitAccountCmds() []*cobra.Command {
	accountCommissionRateCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol (required)")
	accountCommissionRateCmd.MarkFlagRequired("symbol")

	accountIncomeCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol")
	accountIncomeCmd.Flags().StringP("type", "t", "", "Income type")
	accountIncomeCmd.Flags().Int64P("startTime", "a", 0, "Start time (timestamp in ms)")
	accountIncomeCmd.Flags().Int64P("endTime", "e", 0, "End time (timestamp in ms)")
	accountIncomeCmd.Flags().IntP("limit", "l", 100, "Number of results (default 100, max 1000)")

	accountMultiAssetsModeSetCmd.Flags().BoolVar(&multiAssetsModeEnable, "enable", false, "Enable multi-assets mode")
	accountMultiAssetsModeSetCmd.Flags().BoolVar(&multiAssetsModeDisable, "disable", false, "Disable multi-assets mode")
	accountMultiAssetsModeSetCmd.MarkFlagsMutuallyExclusive("enable", "disable")
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

func balances(cmd *cobra.Command, args []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	balances, err := client.GetBalances()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(balances)
}

func accountInfo(cmd *cobra.Command, args []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	account, err := client.GetAccount()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&futures.AccountInfo{AccountResponse: account})
}

func showAccountCommissionRate(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	commissionRate, err := client.GetCommissionRate(symbol)
	if err != nil {
		log.Fatalf("futures commission rate error: %v", err)
	}
	printer.PrintTable(&commissionRate)
}

func showAccountIncome(cmd *cobra.Command, _ []string) {
	symbol, _ := cmd.Flags().GetString("symbol")
	incomeType, _ := cmd.Flags().GetString("type")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	limit, _ := cmd.Flags().GetInt("limit")

	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	income, err := client.GetIncome(symbol, incomeType, startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&income)
}

func showAccountMultiAssetsMode(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	multiAssetsMode, err := client.GetMultiAssetsMode()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("multi-assets mode is:", multiAssetsMode)
}

func setAccountMultiAssetsMode(cmd *cobra.Command, _ []string) {
	if !multiAssetsModeEnable && !multiAssetsModeDisable {
		log.Fatal("Please specify --enable or --disable")
	}

	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	err := client.SetMultiAssetsMode(multiAssetsModeEnable)
	if err != nil {
		log.Fatal(err)
	}

	if multiAssetsModeEnable {
		fmt.Println("Multi-assets mode enabled")
	} else {
		fmt.Println("Multi-assets mode disabled")
	}
}
