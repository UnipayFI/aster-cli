package futures

import (
	"context"
	"time"
)

func (c *Client) GetFundingRate(symbol string, startTime, endTime time.Time, limit int) (FundingRateList, error) {
	service := c.futuresClient().NewGetFundingRateService()
	if symbol != "" {
		service.SetSymbol(symbol)
	}
	if !startTime.IsZero() {
		service.SetStartTime(startTime)
	}
	if !endTime.IsZero() {
		service.SetEndTime(endTime)
	}
	if limit != 0 {
		service.SetLimit(limit)
	}
	rates, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return rates, nil
}

func (c *Client) GetFundingInfo(symbol string) (FundingInfoList, error) {
	service := c.futuresClient().NewGetFundingInfoService()
	if symbol != "" {
		service.SetSymbol(symbol)
	}
	info, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return info, nil
}
