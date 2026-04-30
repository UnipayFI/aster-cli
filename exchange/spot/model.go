package spot

import (
	"fmt"

	"github.com/UnipayFI/aster-cli/common"
	"github.com/UnipayFI/aster-cli/printer"
	"github.com/UnipayFI/go-aster/v3/spot"
)

var _ printer.TableWriter = (*Account)(nil)

type Account struct {
	spot.AccountInfo
}

func (a *Account) Header() []string {
	return []string{"Fee Tier", "Can Trade", "Can Withdraw", "Can Deposit", "Update Time"}
}

func (a *Account) Row() [][]any {
	return [][]any{
		{a.FeeTier, a.CanTrade, a.CanWithdraw, a.CanDeposit, common.FormatUnixTime(a.UpdateTime)},
	}
}

var _ printer.TableWriter = (*AssetBalanceList)(nil)

type AssetBalanceList []spot.Balance

func (a *AssetBalanceList) Header() []string {
	return []string{"Asset", "Free", "Locked"}
}

func (a *AssetBalanceList) Row() [][]any {
	rows := [][]any{}
	for _, asset := range *a {
		if asset.Free.IsZero() && asset.Locked.IsZero() {
			continue
		}
		rows = append(rows, []any{asset.Asset, asset.Free, asset.Locked})
	}
	return rows
}

var _ printer.TableWriter = (*OrderList)(nil)

func (o *OrderList) Header() []string {
	return []string{"Order ID", "Client Order ID", "Symbol", "Side", "Type", "Status", "Price", "Avg Price", "Quantity", "Executed Qty", "Cum Quote", "TIF", "Time", "Update Time"}
}

func (o *OrderList) Row() [][]any {
	rows := [][]any{}
	for _, order := range *o {
		price := order.Price
		if order.Type == spot.OrderTypeMarket && price.IsZero() && !order.ExecutedQty.IsZero() {
			price = order.CumQuote.Div(order.ExecutedQty)
		}
		// Aster's Place/Cancel responses populate transactTime, not the SDK's
		// `time` field, so order.Time is zero on those code paths. Fall back
		// to UpdateTime, which the same responses do populate.
		timeStr := common.FormatTime(order.Time)
		if timeStr == "" {
			timeStr = common.FormatTime(order.UpdateTime)
		}
		rows = append(rows, []any{order.OrderId, order.ClientOrderId, order.Symbol, order.Side, order.Type, order.Status, price, order.AvgPrice, order.OrigQty, order.ExecutedQty, order.CumQuote, order.TimeInForce, timeStr, common.FormatTime(order.UpdateTime)})
	}
	return rows
}

var _ printer.TableWriter = (*TradeList)(nil)

func (t *TradeList) Header() []string {
	return []string{"ID", "Order ID", "Symbol", "Side", "Price", "Quantity", "Quote Quantity", "Commission", "Time", "Maker"}
}

func (t *TradeList) Row() [][]any {
	rows := [][]any{}
	for _, trade := range *t {
		commission := fmt.Sprintf("%s %s", trade.Commission, trade.CommissionAsset)
		rows = append(rows, []any{trade.ID, trade.OrderID, trade.Symbol, trade.Side, trade.Price, trade.Qty, trade.QuoteQty, commission, common.FormatTime(trade.Time), trade.Maker})
	}
	return rows
}

func FilterNonZeroBalances(balances []spot.Balance) *AssetBalanceList {
	list := AssetBalanceList{}
	for _, b := range balances {
		if !b.Free.IsZero() || !b.Locked.IsZero() {
			list = append(list, b)
		}
	}
	return &list
}

var _ printer.TableWriter = (*CommissionRate)(nil)

func (c *CommissionRate) Header() []string {
	return []string{"Symbol", "Maker Commission Rate", "Taker Commission Rate"}
}

func (c *CommissionRate) Row() [][]any {
	return [][]any{
		{c.Symbol, c.MakerCommissionRate, c.TakerCommissionRate},
	}
}
