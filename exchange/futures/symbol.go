package futures

import (
	"context"

	"github.com/UnipayFI/go-aster/futures"
)

func (c *Client) SetLeverage(symbol string, leverage int) (*futures.ChangeLeverageResponse, error) {
	return c.futuresClient().NewChangeLeverageService(symbol, leverage).Do(context.Background())
}

func (c *Client) SetMarginType(symbol string, marginType futures.MarginType) error {
	return c.futuresClient().NewChangeMarginTypeService(symbol, marginType).Do(context.Background())
}

func (c *Client) GetLeverageBrackets(symbol string) (*LeverageBracketList, error) {
	service := c.futuresClient().NewGetLeverageBracketService()
	var brackets []futures.LeverageBracketResponse
	var err error
	if symbol != "" {
		brackets, err = service.Do(context.Background(), symbol)
	} else {
		brackets, err = service.DoAll(context.Background())
	}
	if err != nil {
		return nil, err
	}
	result := LeverageBracketList(brackets)
	return &result, nil
}
