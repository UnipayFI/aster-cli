package wallet

import (
	"context"

	"github.com/UnipayFI/go-aster/futures"
	"github.com/google/uuid"
)

func (c *Client) Transfer(transferType, asset string, amount float64) (*futures.WalletTransferResponse, error) {
	futuresClient := c.NewFuturesClient()
	clientTranId := uuid.New().String()
	return futuresClient.NewWalletTransferService(asset, amount, clientTranId, futures.TransferType(transferType)).Do(context.Background())
}

func (c *Client) TransferWithClientId(transferType, asset string, amount float64, clientTranId string) (*futures.WalletTransferResponse, error) {
	futuresClient := c.NewFuturesClient()
	return futuresClient.NewWalletTransferService(asset, amount, clientTranId, futures.TransferType(transferType)).Do(context.Background())
}
