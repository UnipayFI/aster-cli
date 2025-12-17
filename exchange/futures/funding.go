package futures

import (
	"context"
	"time"
)

func (c *Client) GetFundingRate(symbol string, startTime, endTime int64, limit int) (FundingRateList, error) {
	service := c.futuresClient().NewFundingRateService()
	if symbol != "" {
		service.SetSymbol(symbol)
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
	rates, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return rates, nil
}

func (c *Client) GetFundingInfo(symbol string) (FundingInfoList, error) {
	service := c.futuresClient().NewFundingInfoService()
	if symbol != "" {
		service.SetSymbol(symbol)
	}
	info, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return info, nil
}
