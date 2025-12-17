package exchange

import (
	"github.com/UnipayFI/go-aster/client"
	"github.com/UnipayFI/go-aster/futures"
	"github.com/UnipayFI/go-aster/spot"
)

type Client struct {
	ApiKey    string
	ApiSecret string
}

func NewClient(apiKey, apiSecret string) *Client {
	return &Client{ApiKey: apiKey, ApiSecret: apiSecret}
}

func (c *Client) NewSpotClient() *spot.SpotClient {
	return spot.NewSpotClient(client.WithAuth(c.ApiKey, c.ApiSecret))
}

func (c *Client) NewFuturesClient() *futures.FuturesClient {
	return futures.NewFuturesClient(client.WithAuth(c.ApiKey, c.ApiSecret))
}
