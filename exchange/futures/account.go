package futures

import (
	"context"
	"time"

	"github.com/UnipayFI/aster-cli/exchange"
	"github.com/UnipayFI/go-aster/futures"
)

type Client struct {
	*exchange.Client
}

func (c *Client) futuresClient() *futures.FuturesClient {
	return c.NewFuturesClient()
}

func (c *Client) GetBalances() (*BalanceList, error) {
	balances, err := c.futuresClient().NewGetBalanceService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	result := &BalanceList{}
	for i := range balances {
		if balances[i].Balance.IsZero() && balances[i].CrossWalletBalance.IsZero() && balances[i].CrossUnPnl.IsZero() && balances[i].AvailableBalance.IsZero() {
			continue
		}
		*result = append(*result, balances[i])
	}
	return result, nil
}

func (c *Client) GetAccount() (*futures.AccountResponse, error) {
	return c.futuresClient().NewGetAccountService().Do(context.Background())
}

func (c *Client) GetCommissionRate(symbol string) (CommissionRateList, error) {
	commissionRate, err := c.futuresClient().NewGetCommissionRateService(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return CommissionRateList{*commissionRate}, nil
}

func (c *Client) GetIncome(symbol string, incomeType string, startTime int64, endTime int64, limit int) (IncomeHistoryList, error) {
	service := c.futuresClient().NewGetIncomeService()
	if symbol != "" {
		service.SetSymbol(symbol)
	}
	if incomeType != "" {
		service.SetIncomeType(futures.IncomeType(incomeType))
	}
	if startTime != 0 {
		service.SetStartTime(time.UnixMilli(startTime))
	}
	if endTime != 0 {
		service.SetEndTime(time.UnixMilli(endTime))
	}
	if limit != 0 {
		service.SetLimit(limit)
	}
	income, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return IncomeHistoryList(income), nil
}

func (c *Client) GetMultiAssetsMode() (bool, error) {
	return c.futuresClient().NewGetMultiAssetsModeService().Do(context.Background())
}

func (c *Client) SetMultiAssetsMode(multiAssetsMode bool) error {
	return c.futuresClient().NewChangeMultiAssetsModeService(multiAssetsMode).Do(context.Background())
}
