package spot

import (
	"context"

	"github.com/UnipayFI/aster-cli/exchange"
	"github.com/UnipayFI/go-aster/spot"
)

type Client struct {
	*exchange.Client
}

func (c *Client) GetAccountInfo() (*Account, error) {
	account, err := c.NewSpotClient().NewGetAccountService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &Account{AccountResponse: *account}, nil
}

func (c *Client) GetUserTrades(symbol string, orderId, fromId int64, limit int) (*TradeList, error) {
	service := c.NewSpotClient().NewGetUserTradesService().SetSymbol(symbol)
	if orderId != 0 {
		service.SetOrderId(orderId)
	}
	if fromId != 0 {
		service.SetFromId(fromId)
	}
	if limit != 0 {
		service.SetLimit(limit)
	}
	trades, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := TradeList(trades)
	return &list, nil
}

type TradeList []spot.UserTradeResponse
