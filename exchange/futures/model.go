package futures

import (
	"fmt"

	"github.com/UnipayFI/aster-cli/common"
	"github.com/UnipayFI/aster-cli/printer"
	"github.com/UnipayFI/go-aster/v3/futures"
)

var _ printer.TableWriter = (*BalanceList)(nil)
var _ printer.TableWriter = (*OrderList)(nil)
var _ printer.TableWriter = (*TradeList)(nil)

type BalanceList []futures.Balance

func (a *BalanceList) Header() []string {
	return []string{"Asset", "Balance", "Cross Wallet Balance", "Cross Un Pnl", "Available Balance", "Max Withdraw Amount"}
}

func (a *BalanceList) Row() [][]any {
	rows := [][]any{}
	for _, asset := range *a {
		rows = append(rows, []any{asset.Asset, asset.Balance, asset.CrossWalletBalance, asset.CrossUnPnl, asset.AvailableBalance, asset.MaxWithdrawAmount})
	}
	return rows
}

type AccountInfo struct {
	*futures.AccountInfo
}

func (a *AccountInfo) Header() []string {
	return []string{"Fee Tier", "Can Trade", "Can Deposit", "Can Withdraw", "Total Wallet Balance", "Available Balance"}
}

func (a *AccountInfo) Row() [][]any {
	return [][]any{{a.FeeTier, a.CanTrade, a.CanDeposit, a.CanWithdraw, a.TotalWalletBalance, a.AvailableBalance}}
}

type ForceOrderList []futures.Order

func (f *ForceOrderList) Header() []string {
	return []string{"Order ID", "Symbol", "Side", "Position Side", "Status", "Price", "Quantity", "Executed Quantity", "Time", "Update Time"}
}

func (f *ForceOrderList) Row() [][]any {
	rows := [][]any{}
	for _, order := range *f {
		rows = append(rows, []any{order.OrderId, order.Symbol, order.Side, order.PositionSide, order.Status, order.Price, order.OrigQty, order.ExecutedQty, common.FormatTime(order.Time), common.FormatTime(order.UpdateTime)})
	}
	return rows
}

type OrderList []futures.Order

func (o *OrderList) Header() []string {
	return []string{"Order ID", "Symbol", "Side", "Type", "Position Side", "Status", "Price", "Avg Price", "Quantity", "Executed Qty", "Cum Quote", "TIF", "Time", "Update Time"}
}

func (o *OrderList) Row() [][]any {
	rows := [][]any{}
	for _, order := range *o {
		// See spot OrderList note: Place/Cancel responses populate transactTime
		// (which the SDK doesn't capture), leaving order.Time zero. Fall back
		// to UpdateTime so the table shows a meaningful timestamp.
		timeStr := common.FormatTime(order.Time)
		if timeStr == "" {
			timeStr = common.FormatTime(order.UpdateTime)
		}
		rows = append(rows, []any{order.OrderId, order.Symbol, order.Side, order.Type, order.PositionSide, order.Status, order.Price, order.AvgPrice, order.OrigQty, order.ExecutedQty, order.CumQuote, order.TimeInForce, timeStr, common.FormatTime(order.UpdateTime)})
	}
	return rows
}

type PositionList []futures.AccountPosition

func (p *PositionList) Header() []string {
	return []string{"Symbol", "Position Side", "Position Amount", "Entry Price", "Unrealized Profit", "Leverage", "Update Time"}
}

func (p *PositionList) Row() [][]any {
	rows := [][]any{}
	for _, position := range *p {
		rows = append(rows, []any{position.Symbol, position.PositionSide, position.PositionAmt, position.EntryPrice, position.UnrealizedProfit, position.Leverage, common.FormatUnixTime(position.UpdateTime)})
	}
	return rows
}

type IncomeHistoryList []futures.IncomeRecord

func (i *IncomeHistoryList) Header() []string {
	return []string{"Asset", "Income", "Income Type", "Info", "Symbol", "Time", "Tran ID", "Trade ID"}
}

func (i *IncomeHistoryList) Row() [][]any {
	rows := [][]any{}
	for _, income := range *i {
		rows = append(rows, []any{income.Asset, income.Income, income.IncomeType, income.Info, income.Symbol, common.FormatTime(income.Time), income.TranID, income.TradeID})
	}
	return rows
}

type TradeList []futures.UserTrade

func (t *TradeList) Header() []string {
	return []string{"Trade ID", "Order ID", "Symbol", "Side", "Position Side", "Price", "Quantity", "Quote Quantity", "Realized Pnl", "Commission", "Maker", "Time"}
}

func (t *TradeList) Row() [][]any {
	rows := [][]any{}
	for _, trade := range *t {
		commission := fmt.Sprintf("%s %s", trade.Commission, trade.CommissionAsset)
		rows = append(rows, []any{trade.ID, trade.OrderID, trade.Symbol, trade.Side, trade.PositionSide, trade.Price, trade.Qty, trade.QuoteQty, trade.RealizedPnl, commission, trade.Maker, common.FormatTime(trade.Time)})
	}
	return rows
}

type PositionRiskList []futures.PositionRisk

func (p *PositionRiskList) Header() []string {
	return []string{"Symbol", "Position Side", "Position Amount", "Entry Price", "Mark Price", "Unrealized Profit", "Liquidation Price", "Leverage", "Update Time"}
}

func (p *PositionRiskList) Row() [][]any {
	rows := [][]any{}
	for _, risk := range *p {
		rows = append(rows, []any{risk.Symbol, risk.PositionSide, risk.PositionAmt, risk.EntryPrice, risk.MarkPrice, risk.UnRealizedProfit, risk.LiquidationPrice, risk.Leverage, common.FormatTime(risk.UpdateTime)})
	}
	return rows
}

type CommissionRateList []futures.CommissionRateResponse

func (c *CommissionRateList) Header() []string {
	return []string{"Symbol", "Maker Commission Rate", "Taker Commission Rate"}
}

func (c *CommissionRateList) Row() [][]any {
	rows := [][]any{}
	for _, rate := range *c {
		rows = append(rows, []any{rate.Symbol, rate.MakerCommissionRate, rate.TakerCommissionRate})
	}
	return rows
}

type FundingRateList []futures.FundingRate

func (f *FundingRateList) Header() []string {
	return []string{"Symbol", "Funding Rate", "Funding Time"}
}

func (f *FundingRateList) Row() [][]any {
	rows := [][]any{}
	for _, rate := range *f {
		rows = append(rows, []any{rate.Symbol, rate.FundingRate, common.FormatTime(rate.FundingTime)})
	}
	return rows
}

type FundingInfoList []futures.FundingInfo

func (f *FundingInfoList) Header() []string {
	return []string{"Symbol", "Interest Rate", "Funding Interval (Hours)", "Funding Fee Cap", "Funding Fee Floor"}
}

func (f *FundingInfoList) Row() [][]any {
	rows := [][]any{}
	for _, info := range *f {
		rows = append(rows, []any{info.Symbol, info.InterestRate, info.FundingIntervalHours, info.FundingFeeCap, info.FundingFeeFloor})
	}
	return rows
}

type LeverageBracketList []futures.LeverageBracket

var _ printer.TableWriter = (*LeverageBracketList)(nil)

func (l *LeverageBracketList) Header() []string {
	return []string{"Symbol", "Bracket", "Initial Leverage", "Notional Cap", "Notional Floor", "Maint Margin Ratio"}
}

func (l *LeverageBracketList) Row() [][]any {
	rows := [][]any{}
	for _, item := range *l {
		for _, bracket := range item.Brackets {
			rows = append(rows, []any{item.Symbol, bracket.Bracket, bracket.InitialLeverage, bracket.NotionalCap, bracket.NotionalFloor, bracket.MaintMarginRatio})
		}
	}
	return rows
}

type AdlQuantileList []futures.ADLQuantile

var _ printer.TableWriter = (*AdlQuantileList)(nil)

func (a *AdlQuantileList) Header() []string {
	return []string{"Symbol", "LONG", "SHORT", "BOTH", "HEDGE"}
}

func (a *AdlQuantileList) Row() [][]any {
	rows := [][]any{}
	for _, item := range *a {
		rows = append(rows, []any{item.Symbol, item.AdlQuantile["LONG"], item.AdlQuantile["SHORT"], item.AdlQuantile["BOTH"], item.AdlQuantile["HEDGE"]})
	}
	return rows
}

type PositionMarginHistoryList []futures.PositionMarginHistoryEntry

var _ printer.TableWriter = (*PositionMarginHistoryList)(nil)

func (p *PositionMarginHistoryList) Header() []string {
	return []string{"Symbol", "Position Side", "Amount", "Asset", "Type", "Time"}
}

func (p *PositionMarginHistoryList) Row() [][]any {
	rows := [][]any{}
	for _, item := range *p {
		marginType := "Add"
		if item.Type == int(futures.PositionMarginReduce) {
			marginType = "Reduce"
		}
		rows = append(rows, []any{item.Symbol, item.PositionSide, item.Amount, item.Asset, marginType, common.FormatTime(item.Time)})
	}
	return rows
}
