package spot

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/UnipayFI/go-aster/spot"
)

func (c *Client) GetOrderList(symbol string, orderID int64, start, end time.Time, limit int) (*OrderList, error) {
	service := c.NewSpotClient().NewGetAllOrdersService(symbol)
	if orderID != 0 {
		service.SetOrderId(orderID)
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
	list := OrderList(orders)
	return &list, nil
}

func (c *Client) GetOpenOrders(symbol string) (*OrderList, error) {
	service := c.NewSpotClient().NewGetOpenOrdersService()
	if symbol != "" {
		service.SetSymbol(symbol)
	}
	orders, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := OrderList(orders)
	return &list, nil
}

func (c *Client) GetOrder(symbol string, orderID int64, clientOrderID string) (*spot.OrderResponse, error) {
	service := c.NewSpotClient().NewGetOrderService(symbol)
	if orderID != 0 {
		service.SetOrderId(orderID)
	}
	if clientOrderID != "" {
		service.SetOrigClientOrderId(clientOrderID)
	}
	return service.Do(context.Background())
}

func (c *Client) CreateOrder(params map[string]string) (*spot.OrderResponse, error) {
	sideType := spot.OrderSide(strings.ToUpper(params["side"]))
	orderType := spot.OrderType(strings.ToUpper(params["type"]))
	orderService := c.NewSpotClient().NewCreateOrderService(params["symbol"], sideType, orderType)

	if params["quantity"] != "" {
		qty, _ := strconv.ParseFloat(params["quantity"], 64)
		orderService.SetQuantity(qty)
	}
	if params["quoteOrderQty"] != "" {
		qty, _ := strconv.ParseFloat(params["quoteOrderQty"], 64)
		orderService.SetQuoteOrderQty(qty)
	}
	if params["timeInForce"] != "" {
		orderService.SetTimeInForce(spot.TimeInForce(strings.ToUpper(params["timeInForce"])))
	} else if orderType == spot.OrderTypeLimit {
		orderService.SetTimeInForce(spot.TimeInForceTypeGTC)
	}
	if params["price"] != "" {
		price, _ := strconv.ParseFloat(params["price"], 64)
		orderService.SetPrice(price)
	}
	if params["newClientOrderId"] != "" {
		orderService.SetNewClientOrderId(params["newClientOrderId"])
	}
	if params["stopPrice"] != "" {
		stopPrice, _ := strconv.ParseFloat(params["stopPrice"], 64)
		orderService.SetStopPrice(stopPrice)
	}

	return orderService.Do(context.Background())
}

func (c *Client) CancelOrder(symbol string, orderID int64, clientOrderID string) error {
	orderService := c.NewSpotClient().NewCancelOrderService(symbol)
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
	orderService := c.NewSpotClient().NewCancelAllOpenOrdersService(symbol)
	_, err := orderService.Do(context.Background())
	return err
}

type OrderList []spot.OrderResponse
