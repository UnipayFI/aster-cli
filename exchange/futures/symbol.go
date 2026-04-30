package futures

import (
	"context"

	"github.com/UnipayFI/go-aster/v3/futures"
)

func (c *Client) SetLeverage(symbol string, leverage int) (*futures.ChangeLeverageResponse, error) {
	return c.futuresClient().NewChangeLeverageService(symbol, leverage).Do(context.Background())
}

func (c *Client) SetMarginType(symbol string, marginType futures.MarginType) (*futures.GenericCodeMsg, error) {
	return c.futuresClient().NewChangeMarginTypeService(symbol, marginType).Do(context.Background())
}

func (c *Client) GetLeverageBrackets(symbol string) (*LeverageBracketList, error) {
	service := c.futuresClient().NewGetLeverageBracketService()
	if symbol != "" {
		service.SetSymbol(symbol)
	}
	brackets, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	result := LeverageBracketList(brackets)
	return &result, nil
}
