package futures

import (
	"log"

	"github.com/UnipayFI/aster-cli/common"
	"github.com/UnipayFI/aster-cli/config"
	"github.com/UnipayFI/aster-cli/exchange"
	"github.com/UnipayFI/aster-cli/exchange/futures"
	"github.com/UnipayFI/aster-cli/printer"
	"github.com/spf13/cobra"
)

var (
	fundingCmd = &cobra.Command{
		Use:   "funding",
		Short: "Funding rate commands",
		Long:  `Query funding rate history and funding info.`,
	}

	// funding info
	fundingInfoCmd = &cobra.Command{
		Use:     "info",
		Aliases: []string{"i"},
		Short:   "Query funding info",
		Long:    `Query funding info including funding interval and fee cap/floor.`,
		Run:     showFundingInfo,
	}

	// funding rate
	fundingRateCmd = &cobra.Command{
		Use:     "rate",
		Aliases: []string{"r"},
		Short:   "Query funding rate history",
		Long:    `Query funding rate history.`,
		Run:     showFundingRate,
	}
)

func InitFundingCmds() []*cobra.Command {
	fundingInfoCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol")

	fundingRateCmd.Flags().StringP("symbol", "s", "", "Trading pair symbol")
	fundingRateCmd.Flags().StringP("startTime", "a", "", "Start time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	fundingRateCmd.Flags().StringP("endTime", "e", "", "End time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	fundingRateCmd.Flags().IntP("limit", "l", 100, "Number of results (default 100, max 1000)")

	fundingCmd.AddCommand(fundingInfoCmd, fundingRateCmd)
	return []*cobra.Command{fundingCmd}
}

func showFundingInfo(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	info, err := client.GetFundingInfo(symbol)
	if err != nil {
		log.Fatalf("futures funding info error: %v", err)
	}
	printer.PrintTable(&info)
}

func showFundingRate(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
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
	rates, err := client.GetFundingRate(symbol, startTime, endTime, limit)
	if err != nil {
		log.Fatalf("futures funding rate error: %v", err)
	}
	printer.PrintTable(&rates)
}
