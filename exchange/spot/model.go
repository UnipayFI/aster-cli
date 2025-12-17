package spot

import (
	"fmt"

	"github.com/UnipayFI/aster-cli/printer"
	"github.com/UnipayFI/go-aster/spot"
)

var _ printer.TableWriter = (*Account)(nil)

type Account struct {
	spot.AccountResponse
}

func (a *Account) Header() []string {
	return []string{"Fee Tier", "Can Trade", "Can Withdraw", "Can Deposit", "Update Time"}
}

func (a *Account) Row() [][]any {
	return [][]any{
		{a.FeeTier, a.CanTrade, a.CanWithdraw, a.CanDeposit, a.UpdateTime.Format("2006-01-02 15:04:05")},
	}
}

var _ printer.TableWriter = (*AssetBalanceList)(nil)

type AssetBalanceList []spot.AccountBalance

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
		rows = append(rows, []any{order.OrderId, order.ClientOrderId, order.Symbol, order.Side, order.Type, order.Status, price, order.AvgPrice, order.OrigQty, order.ExecutedQty, order.CumQuote, order.TimeInForce, order.Time.Format("2006-01-02 15:04:05"), order.UpdateTime.Format("2006-01-02 15:04:05")})
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
		rows = append(rows, []any{trade.Id, trade.OrderId, trade.Symbol, trade.Side, trade.Price, trade.Qty, trade.QuoteQty, commission, trade.Time.Format("2006-01-02 15:04:05"), trade.Maker})
	}
	return rows
}

func FilterNonZeroBalances(balances []spot.AccountBalance) *AssetBalanceList {
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
