package wallet

import (
	"context"

	"github.com/UnipayFI/go-aster/v3/futures"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (c *Client) Transfer(transferType, asset string, amount decimal.Decimal) (*futures.TransferResponse, error) {
	clientTranId := uuid.New().String()
	return c.TransferWithClientId(transferType, asset, amount, clientTranId)
}

func (c *Client) TransferWithClientId(transferType, asset string, amount decimal.Decimal, clientTranId string) (*futures.TransferResponse, error) {
	futuresClient := c.NewFuturesClient()
	return futuresClient.NewFuturesSpotTransferService(asset, amount, futures.TransferKindType(transferType), clientTranId).Do(context.Background())
}
