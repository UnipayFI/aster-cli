package futures

import (
	"context"
	"strings"
	"time"

	"github.com/UnipayFI/go-aster/v3/futures"
	"github.com/shopspring/decimal"
)

func (c *Client) GetOrderList(symbol string, limit int, start, end time.Time, orderID int64) (OrderList, error) {
	service := c.futuresClient().NewGetAllOrdersService(symbol)
	if limit != 0 {
		service.SetLimit(limit)
	}
	if orderID != 0 {
		service.SetOrderId(orderID)
	}
	if !start.IsZero() {
		service.SetStartTime(start)
	}
	if !end.IsZero() {
		service.SetEndTime(end)
	}
	orders, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (c *Client) GetOpenOrders(symbol string) (OrderList, error) {
	service := c.futuresClient().NewGetOpenOrdersService()
	if symbol != "" {
		service.SetSymbol(symbol)
	}
	orders, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (c *Client) GetOrder(symbol string, orderID int64, clientOrderID string) (*futures.Order, error) {
	service := c.futuresClient().NewGetOrderService(symbol)
	if orderID != 0 {
		service.SetOrderId(orderID)
	}
	if clientOrderID != "" {
		service.SetOrigClientOrderId(clientOrderID)
	}
	return service.Do(context.Background())
}

func (c *Client) GetForceOrders(symbol string, autoCloseType futures.AutoCloseType, start, end time.Time, limit int) (ForceOrderList, error) {
	service := c.futuresClient().NewGetForceOrdersService()
	if symbol != "" {
		service.SetSymbol(symbol)
	}
	if autoCloseType != "" {
		service.SetAutoCloseType(autoCloseType)
	}
	if !start.IsZero() {
		service.SetStartTime(start)
	}
	if !end.IsZero() {
		service.SetEndTime(end)
	}
	if limit != 0 {
		service.SetLimit(limit)
	}
	orders, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return ForceOrderList(orders), nil
}

func (c *Client) CreateOrder(params map[string]string) (*futures.Order, error) {
	side := futures.OrderSide(strings.ToUpper(params["side"]))
	orderType := futures.OrderType(strings.ToUpper(params["type"]))
	service := c.futuresClient().NewPlaceOrderService(params["symbol"], side, orderType)

	if v := params["positionSide"]; v != "" {
		service.SetPositionSide(futures.PositionSide(strings.ToUpper(v)))
	}
	if v := params["quantity"]; v != "" {
		q, err := decimal.NewFromString(v)
		if err != nil {
			return nil, err
		}
		service.SetQuantity(q)
	}
	if v := params["reduceOnly"]; v != "" {
		service.SetReduceOnly(v == "true")
	}
	// timeInForce is required for LIMIT orders, default to GTC if not provided
	if v := params["timeInForce"]; v != "" {
		service.SetTimeInForce(futures.TimeInForce(strings.ToUpper(v)))
	} else if orderType == futures.OrderTypeLimit {
		service.SetTimeInForce(futures.TimeInForceGTC)
	}
	if v := params["price"]; v != "" {
		p, err := decimal.NewFromString(v)
		if err != nil {
			return nil, err
		}
		service.SetPrice(p)
	}
	if v := params["newClientOrderId"]; v != "" {
		service.SetNewClientOrderId(v)
	}
	if v := params["stopPrice"]; v != "" {
		p, err := decimal.NewFromString(v)
		if err != nil {
			return nil, err
		}
		service.SetStopPrice(p)
	}
	if v := params["closePosition"]; v != "" {
		service.SetClosePosition(v == "true")
	}
	if v := params["activationPrice"]; v != "" {
		p, err := decimal.NewFromString(v)
		if err != nil {
			return nil, err
		}
		service.SetActivationPrice(p)
	}
	if v := params["callbackRate"]; v != "" {
		r, err := decimal.NewFromString(v)
		if err != nil {
			return nil, err
		}
		service.SetCallbackRate(r)
	}
	if v := params["workingType"]; v != "" {
		service.SetWorkingType(futures.WorkingType(strings.ToUpper(v)))
	}
	if v := params["priceProtect"]; v != "" {
		service.SetPriceProtect(v == "true")
	}
	if v := params["newOrderRespType"]; v != "" {
		service.SetNewOrderRespType(futures.ResponseType(strings.ToUpper(v)))
	}

	return service.Do(context.Background())
}

func (c *Client) CancelOrder(symbol string, orderID int64, clientOrderID string) (*futures.Order, error) {
	service := c.futuresClient().NewCancelOrderService(symbol)
	if orderID != 0 {
		service.SetOrderId(orderID)
	}
	if clientOrderID != "" {
		service.SetOrigClientOrderId(clientOrderID)
	}
	return service.Do(context.Background())
}

func (c *Client) CancelAllOrders(symbol string) (*futures.GenericCodeMsg, error) {
	return c.futuresClient().NewCancelAllOpenOrdersService(symbol).Do(context.Background())
}

func (c *Client) GetTrades(symbol string, start, end time.Time, fromId int64, limit int) (TradeList, error) {
	service := c.futuresClient().NewGetUserTradesService(symbol)
	if !start.IsZero() {
		service.SetStartTime(start)
	}
	if !end.IsZero() {
		service.SetEndTime(end)
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
	return trades, nil
}
