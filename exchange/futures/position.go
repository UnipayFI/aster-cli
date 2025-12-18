package futures

import (
	"context"
	"slices"
	"time"

	"github.com/UnipayFI/go-aster/futures"
	"github.com/shopspring/decimal"
)

func (c *Client) GetPositions() (PositionList, error) {
	account, err := c.futuresClient().NewGetAccountService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	positions := PositionList{}
	for _, position := range account.Positions {
		if !position.PositionAmt.IsZero() {
			positions = append(positions, position)
		}
	}
	return positions, nil
}

func (c *Client) GetPositionRisk(symbol string) (PositionRiskList, error) {
	service := c.futuresClient().NewGetPositionRiskService()
	if symbol != "" {
		service.SetSymbol(symbol)
	}
	risks, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	risks = slices.DeleteFunc(risks, func(risk futures.PositionRiskResponse) bool {
		return risk.PositionAmt.IsZero()
	})
	return risks, nil
}

func (c *Client) ModifyPositionMargin(symbol, positionSide string, amount float64, typ int) error {
	service := c.futuresClient().NewAddIsolatedMarginService(symbol, decimal.NewFromFloat(amount), typ)
	if positionSide != "" {
		service.SetPositionSide(futures.PositionSide(positionSide))
	}
	_, err := service.Do(context.Background())
	return err
}

func (c *Client) GetPositionSide() (bool, error) {
	return c.futuresClient().NewGetPositionModeService().Do(context.Background())
}

func (c *Client) ChangePositionSide(dualSide bool) error {
	return c.futuresClient().NewChangePositionModeService(dualSide).Do(context.Background())
}

func (c *Client) GetPositionMode() (bool, error) {
	return c.futuresClient().NewGetPositionModeService().Do(context.Background())
}

func (c *Client) ChangePositionMode(dualSidePosition bool) error {
	return c.futuresClient().NewChangePositionModeService(dualSidePosition).Do(context.Background())
}

func (c *Client) GetPositionMarginHistory(symbol string, marginType int, startTime, endTime time.Time, limit int) (*PositionMarginHistoryList, error) {
	service := c.futuresClient().NewGetPositionMarginHistoryService(symbol)
	if marginType != 0 {
		service.SetType(marginType)
	}
	if !startTime.IsZero() {
		service.SetStartTime(startTime)
	}
	if !endTime.IsZero() {
		service.SetEndTime(endTime)
	}
	if limit > 0 {
		service.SetLimit(limit)
	}
	history, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	result := PositionMarginHistoryList(history)
	return &result, nil
}
func (c *Client) GetAdlQuantile(symbol string) (*AdlQuantileList, error) {
	service := c.futuresClient().NewGetAdlQuantileService()
	var quantiles []futures.AdlQuantileResponse
	var err error
	if symbol != "" {
		var resp *futures.AdlQuantileResponse
		resp, err = service.Do(context.Background(), symbol)
		if err != nil {
			return nil, err
		}
		quantiles = []futures.AdlQuantileResponse{*resp}
	} else {
		quantiles, err = service.DoAll(context.Background())
		if err != nil {
			return nil, err
		}
	}
	result := AdlQuantileList(quantiles)
	return &result, nil
}
