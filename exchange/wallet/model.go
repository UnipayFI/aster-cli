package wallet

import (
	"github.com/UnipayFI/aster-cli/printer"
	"github.com/UnipayFI/go-aster/v3/futures"
)

var _ printer.TableWriter = (*TransferResult)(nil)

type TransferResult struct {
	*futures.TransferResponse
}

func (t *TransferResult) Header() []string {
	return []string{"Transaction ID", "Status"}
}

func (t *TransferResult) Row() [][]any {
	return [][]any{{t.TranID, t.Status}}
}
