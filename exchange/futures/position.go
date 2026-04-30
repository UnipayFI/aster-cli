package futures

import (
	"context"
	"slices"
	"time"

	"github.com/UnipayFI/go-aster/v3/futures"
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
	risks = slices.DeleteFunc(risks, func(risk futures.PositionRisk) bool {
		return risk.PositionAmt.IsZero()
	})
	return risks, nil
}

func (c *Client) ModifyPositionMargin(symbol, positionSide string, amount decimal.Decimal, action futures.PositionMarginType) (*futures.ModifyMarginResponse, error) {
	service := c.futuresClient().NewModifyIsolatedPositionMarginService(symbol, amount, action)
	if positionSide != "" {
		service.SetPositionSide(futures.PositionSide(positionSide))
	}
	return service.Do(context.Background())
}

func (c *Client) GetPositionMode() (bool, error) {
	resp, err := c.futuresClient().NewGetPositionModeService().Do(context.Background())
	if err != nil {
		return false, err
	}
	return resp.DualSidePosition, nil
}

func (c *Client) ChangePositionMode(dualSidePosition bool) (*futures.GenericCodeMsg, error) {
	return c.futuresClient().NewChangePositionModeService(dualSidePosition).Do(context.Background())
}

func (c *Client) GetPositionMarginHistory(symbol string, marginType futures.PositionMarginType, startTime, endTime time.Time, limit int) (*PositionMarginHistoryList, error) {
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
	service := c.futuresClient().NewGetADLQuantileService()
	if symbol != "" {
		service.SetSymbol(symbol)
	}
	quantiles, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	result := AdlQuantileList(quantiles)
	return &result, nil
}
