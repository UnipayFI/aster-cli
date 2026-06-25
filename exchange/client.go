package exchange

import (
	"github.com/UnipayFI/aster-cli/config"
	aster "github.com/UnipayFI/go-aster/v3"
	"github.com/UnipayFI/go-aster/v3/client"
	asterCommon "github.com/UnipayFI/go-aster/v3/common"
	"github.com/UnipayFI/go-aster/v3/futures"
	"github.com/UnipayFI/go-aster/v3/spot"
)

type Client struct {
	Address    string
	PrivateKey string
	ChainID    int64
}

func NewClient(address, privateKey string, chainID int64) *Client {
	return &Client{Address: address, PrivateKey: privateKey, ChainID: chainID}
}

func (c *Client) baseOptions() []client.Options {
	opts := []client.Options{client.WithAuth(c.Address, c.PrivateKey)}
	if c.ChainID == asterCommon.EIP712_CHAIN_ID_TESTNET {
		opts = append(opts, client.WithNetwork(asterCommon.Testnet))
	} else {
		opts = append(opts, client.WithNetwork(asterCommon.Mainnet))
	}
	return opts
}

func (c *Client) NewSpotClient() *spot.SpotClient {
	opts := c.baseOptions()
	if config.Config.SpotBaseURL != "" {
		opts = append(opts, client.WithBaseURL(config.Config.SpotBaseURL))
	}
	return aster.NewSpotClient(opts...)
}

func (c *Client) NewFuturesClient() *futures.FuturesClient {
	opts := c.baseOptions()
	if config.Config.FuturesBaseURL != "" {
		opts = append(opts, client.WithBaseURL(config.Config.FuturesBaseURL))
	}
	return aster.NewFuturesClient(opts...)
}
