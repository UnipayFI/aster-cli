package wallet

import (
	"github.com/UnipayFI/aster-cli/printer"
	"github.com/UnipayFI/go-aster/futures"
)

var _ printer.TableWriter = (*TransferResult)(nil)

type TransferResult struct {
	*futures.WalletTransferResponse
}

func (t *TransferResult) Header() []string {
	return []string{"Transaction ID", "Status"}
}

func (t *TransferResult) Row() [][]any {
	return [][]any{{t.TranId, t.Status}}
}
