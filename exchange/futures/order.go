package futures

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/UnipayFI/go-aster/futures"
)

func (c *Client) GetOrderList(symbol string, limit int, start, end int64, orderID int64) (OrderList, error) {
	service := c.futuresClient().NewGetAllOrdersService(symbol)
	if limit != 0 {
		service.SetLimit(limit)
	}
	if orderID != 0 {
		service.SetOrderId(orderID)
	}
	if start != 0 {
		service.SetStartTime(time.UnixMilli(start))
	}
	if end != 0 {
		service.SetEndTime(time.UnixMilli(end))
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

func (c *Client) GetOrder(symbol string, orderID int64, clientOrderID string) (*futures.OrderResponse, error) {
	service := c.futuresClient().NewGetOrderService(symbol)
	if orderID != 0 {
		service.SetOrderId(orderID)
	}
	if clientOrderID != "" {
		service.SetOrigClientOrderId(clientOrderID)
	}
	return service.Do(context.Background())
}

func (c *Client) GetForceOrders(symbol string, autoCloseType futures.AutoCloseType, startTime, endTime int64, limit int) (ForceOrderList, error) {
	service := c.futuresClient().NewGetForceOrdersService()
	if symbol != "" {
		service.SetSymbol(symbol)
	}
	if autoCloseType != "" {
		service.SetAutoCloseType(autoCloseType)
	}
	if startTime != 0 {
		service.SetStartTime(time.UnixMilli(startTime))
	}
	if endTime != 0 {
		service.SetEndTime(time.UnixMilli(endTime))
	}
	if limit != 0 {
		service.SetLimit(limit)
	}
	orders, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (c *Client) CreateOrder(params map[string]string) (*futures.OrderResponse, error) {
	sideType := futures.OrderSide(strings.ToUpper(params["side"]))
	orderType := futures.OrderType(strings.ToUpper(params["type"]))
	orderService := c.futuresClient().NewCreateOrderService(params["symbol"], sideType, orderType)

	if params["positionSide"] != "" {
		orderService.SetPositionSide(futures.PositionSide(params["positionSide"]))
	}
	if params["quantity"] != "" {
		qty, _ := strconv.ParseFloat(params["quantity"], 64)
		orderService.SetQuantity(qty)
	}
	if params["reduceOnly"] != "" {
		orderService.SetReduceOnly(params["reduceOnly"] == "true")
	}
	// timeInForce is required for LIMIT orders, default to GTC if not provided
	if params["timeInForce"] != "" {
		orderService.SetTimeInForce(futures.TimeInForce(strings.ToUpper(params["timeInForce"])))
	} else if orderType == futures.OrderTypeLimit {
		orderService.SetTimeInForce(futures.TimeInForceGTC)
	}
	if params["price"] != "" {
		price, _ := strconv.ParseFloat(params["price"], 64)
		orderService.SetPrice(price)
	}
	if params["newClientOrderID"] != "" {
		orderService.SetNewClientOrderId(params["newClientOrderID"])
	}
	if params["stopPrice"] != "" {
		stopPrice, _ := strconv.ParseFloat(params["stopPrice"], 64)
		orderService.SetStopPrice(stopPrice)
	}
	if params["closePosition"] != "" {
		orderService.SetClosePosition(params["closePosition"] == "true")
	}
	if params["activationPrice"] != "" {
		activationPrice, _ := strconv.ParseFloat(params["activationPrice"], 64)
		orderService.SetActivationPrice(activationPrice)
	}
	if params["callbackRate"] != "" {
		callbackRate, _ := strconv.ParseFloat(params["callbackRate"], 64)
		orderService.SetCallbackRate(callbackRate)
	}
	if params["workingType"] != "" {
		orderService.SetWorkingType(futures.WorkingType(params["workingType"]))
	}
	if params["priceProtect"] != "" {
		orderService.SetPriceProtect(params["priceProtect"] == "true")
	}
	if params["newOrderRespType"] != "" {
		orderService.SetNewOrderRespType(futures.NewOrderRespType(params["newOrderRespType"]))
	}

	return orderService.Do(context.Background())
}

func (c *Client) CancelOrder(symbol string, orderID int64, clientOrderID string) error {
	orderService := c.futuresClient().NewCancelOrderService(symbol)
	if orderID != 0 {
		orderService.SetOrderId(orderID)
	}
	if clientOrderID != "" {
		orderService.SetOrigClientOrderId(clientOrderID)
	}
	_, err := orderService.Do(context.Background())
	return err
}

func (c *Client) CancelAllOrders(symbol string) error {
	_, err := c.futuresClient().NewCancelAllOrdersService(symbol).Do(context.Background())
	return err
}

func (c *Client) GetTrades(symbol string, startTime, endTime int64, fromId int64, limit int) (TradeList, error) {
	service := c.futuresClient().NewGetUserTradesService(symbol)
	if startTime != 0 {
		service.SetStartTime(time.UnixMilli(startTime))
	}
	if endTime != 0 {
		service.SetEndTime(time.UnixMilli(endTime))
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
