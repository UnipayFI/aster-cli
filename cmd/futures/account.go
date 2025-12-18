package futures

import (
	"fmt"
	"log"

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

	multiAssetsMargin            bool
	accountMultiAssetsModeSetCmd = &cobra.Command{
		Use:   "set",
		Short: "Set multi-assets mode",
		Long:  `Change multi-assets mode. Use --multiAssetsMargin=true for Multi-Assets Mode or --multiAssetsMargin=false for Single-Asset Mode.`,
		Run:   setAccountMultiAssetsMode,
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
	incomeType, _ := cmd.Flags().GetString("incomeType")
	startTimeRaw, _ := cmd.Flags().GetString("startTime")
	endTimeRaw, _ := cmd.Flags().GetString("endTime")
	limit, _ := cmd.Flags().GetInt("limit")

	startTime, _, err := common.ParseTimeFlagUnixMilli("--startTime", startTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	endTime, _, err := common.ParseTimeFlagUnixMilli("--endTime", endTimeRaw)
	if err != nil {
		log.Fatal(err)
	}

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
	if !cmd.Flags().Changed("multiAssetsMargin") {
		log.Fatal("Please specify --multiAssetsMargin")
	}

	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	err := client.SetMultiAssetsMode(multiAssetsMargin)
	if err != nil {
		log.Fatal(err)
	}

	if multiAssetsMargin {
		fmt.Println("Multi-assets mode enabled")
	} else {
		fmt.Println("Multi-assets mode disabled")
	}
}
