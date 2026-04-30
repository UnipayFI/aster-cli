package spot

import (
	"context"
	"strings"
	"time"

	"github.com/UnipayFI/go-aster/v3/spot"
	"github.com/shopspring/decimal"
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
	side := spot.OrderSide(strings.ToUpper(params["side"]))
	orderType := spot.OrderType(strings.ToUpper(params["type"]))
	service := c.NewSpotClient().NewPlaceOrderService(params["symbol"], side, orderType)

	if v := params["quantity"]; v != "" {
		q, err := decimal.NewFromString(v)
		if err != nil {
			return nil, err
		}
		service.SetQuantity(q)
	}
	if v := params["quoteOrderQty"]; v != "" {
		q, err := decimal.NewFromString(v)
		if err != nil {
			return nil, err
		}
		service.SetQuoteOrderQty(q)
	}
	if v := params["timeInForce"]; v != "" {
		service.SetTimeInForce(spot.TimeInForce(strings.ToUpper(v)))
	} else if orderType == spot.OrderTypeLimit {
		service.SetTimeInForce(spot.TimeInForceGTC)
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

	return service.Do(context.Background())
}

func (c *Client) CancelOrder(symbol string, orderID int64, clientOrderID string) (*spot.OrderResponse, error) {
	service := c.NewSpotClient().NewCancelOrderService(symbol)
	if orderID != 0 {
		service.SetOrderId(orderID)
	}
	if clientOrderID != "" {
		service.SetOrigClientOrderId(clientOrderID)
	}
	return service.Do(context.Background())
}

func (c *Client) CancelAllOrders(symbol string) (*spot.CancelAllOpenOrdersResponse, error) {
	return c.NewSpotClient().NewCancelAllOpenOrdersService(symbol).Do(context.Background())
}

type OrderList []spot.OrderResponse
