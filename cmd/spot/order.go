package spot

import (
	"log"
	"time"

	"github.com/UnipayFI/aster-cli/common"
	"github.com/UnipayFI/aster-cli/exchange/spot"
	"github.com/UnipayFI/aster-cli/printer"
	"github.com/spf13/cobra"
)

var (
	orderCmd = &cobra.Command{
		Use:   "order",
		Short: "Support create, cancel, list orders",
	}

	orderListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List orders",
		Long: `Get all account orders.

Docs Link: https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#query-all-orders-user_data`,
		Run: orderList,
	}
	orderOpenListCmd = &cobra.Command{
		Use:   "open",
		Short: "List open orders",
		Long: `Get all open orders on a symbol.

Docs Link: https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#current-open-orders-user_data`,
		Run: orderOpenList,
	}
	orderCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create order",
		Long: `Create a new order.
* Supports parameters: symbol, side, type, quantity, quoteOrderQty, timeInForce, price, newClientOrderId, stopPrice

Docs Link: https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#place-order-trade`,
		Run: createOrder,
	}
	orderCancelCmd = &cobra.Command{
		Use:   "cancel",
		Short: "Cancel order",
		Long: `Cancel order.
If either orderId or origClientOrderId is provided, the specified order will be canceled.
If only the symbol is passed, all open orders for that trading pair will be canceled.

Docs Link: https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#cancel-order-trade`,
		Run: cancelOrder,
	}
	orderGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Query a single order",
		Long: `Query a single order by orderId or origClientOrderId.

Docs Link: https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#query-order-user_data`,
		Run: getOrder,
	}
)

func InitOrderCmds() []*cobra.Command {
	orderCmd.PersistentFlags().StringP("symbol", "s", "", "symbol")

	orderListCmd.Flags().Int64P("orderId", "i", 0, "orderId")
	orderListCmd.Flags().IntP("limit", "l", 500, "limit, max 1000")
	orderListCmd.Flags().StringP("startTime", "a", "", "start time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	orderListCmd.Flags().StringP("endTime", "e", "", "end time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	orderListCmd.MarkFlagRequired("symbol")

	var side, orderType string
	orderCreateCmd.Flags().StringVarP(&side, "side", "S", "", "BUY or SELL")
	orderCreateCmd.Flags().StringVarP(&orderType, "type", "t", "", "LIMIT, MARKET, STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET")
	orderCreateCmd.Flags().StringP("quantity", "q", "", "order quantity (decimal string)")
	orderCreateCmd.Flags().String("quoteOrderQty", "", "quote order quantity for MARKET orders (decimal string)")
	orderCreateCmd.Flags().StringP("price", "p", "", "order price, required for LIMIT orders (decimal string)")
	orderCreateCmd.Flags().StringP("timeInForce", "T", "", "GTC, IOC, FOK (default GTC for LIMIT orders)")
	orderCreateCmd.Flags().String("stopPrice", "", "stop price for STOP/TAKE_PROFIT orders (decimal string)")
	orderCreateCmd.Flags().String("newClientOrderId", "", "custom order id")
	orderCreateCmd.MarkFlagRequired("symbol")

	orderCancelCmd.Flags().Int64P("orderId", "i", 0, "orderId")
	orderCancelCmd.Flags().StringP("origClientOrderId", "c", "", "origClientOrderId")
	orderCancelCmd.MarkFlagRequired("symbol")

	orderGetCmd.Flags().Int64P("orderId", "i", 0, "orderId")
	orderGetCmd.Flags().StringP("origClientOrderId", "c", "", "origClientOrderId")
	orderGetCmd.MarkFlagRequired("symbol")

	orderCmd.AddCommand(orderListCmd, orderOpenListCmd, orderCreateCmd, orderCancelCmd, orderGetCmd)
	return []*cobra.Command{orderCmd}
}

func orderList(cmd *cobra.Command, args []string) {
	client := newClient()
	symbol, _ := cmd.Flags().GetString("symbol")
	limit, _ := cmd.Flags().GetInt("limit")
	startTimeRaw, _ := cmd.Flags().GetString("startTime")
	endTimeRaw, _ := cmd.Flags().GetString("endTime")
	orderID, _ := cmd.Flags().GetInt64("orderId")

	var startTime, endTime time.Time
	if t, ok, err := common.ParseTimeFlag("--startTime", startTimeRaw); err != nil {
		log.Fatal(err)
	} else if ok {
		startTime = t
	}
	if t, ok, err := common.ParseTimeFlag("--endTime", endTimeRaw); err != nil {
		log.Fatal(err)
	} else if ok {
		endTime = t
	}

	orders, err := client.GetOrderList(symbol, orderID, startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(orders)
}

func orderOpenList(cmd *cobra.Command, _ []string) {
	client := newClient()
	symbol, _ := cmd.Flags().GetString("symbol")
	orders, err := client.GetOpenOrders(symbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(orders)
}

func createOrder(cmd *cobra.Command, _ []string) {
	client := newClient()
	order, err := client.CreateOrder(buildOrderParams(cmd))
	if err != nil {
		log.Fatal(err)
	}
	orders := spot.OrderList{*order}
	printer.Print(&orders)
}

// buildOrderParams reads the cobra-parsed flags and returns the long-name
// keyed map that exchange/spot.Client.CreateOrder expects.
func buildOrderParams(cmd *cobra.Command) map[string]string {
	params := map[string]string{}
	stringFlags := []string{
		"symbol", "side", "type", "quantity", "quoteOrderQty",
		"price", "timeInForce", "stopPrice", "newClientOrderId",
	}
	for _, name := range stringFlags {
		if v, _ := cmd.Flags().GetString(name); v != "" {
			params[name] = v
		}
	}
	return params
}


func cancelOrder(cmd *cobra.Command, _ []string) {
	client := newClient()
	symbol, _ := cmd.Flags().GetString("symbol")
	orderID, _ := cmd.Flags().GetInt64("orderId")
	clientOrderID, _ := cmd.Flags().GetString("origClientOrderId")

	if orderID == 0 && clientOrderID == "" {
		resp, err := client.CancelAllOrders(symbol)
		if err != nil {
			log.Fatal(err)
		}
		printer.Print(resp)
		return
	}
	order, err := client.CancelOrder(symbol, orderID, clientOrderID)
	if err != nil {
		log.Fatal(err)
	}
	orders := spot.OrderList{*order}
	printer.Print(&orders)
}

func getOrder(cmd *cobra.Command, _ []string) {
	client := newClient()
	symbol, _ := cmd.Flags().GetString("symbol")
	orderID, _ := cmd.Flags().GetInt64("orderId")
	clientOrderID, _ := cmd.Flags().GetString("origClientOrderId")

	order, err := client.GetOrder(symbol, orderID, clientOrderID)
	if err != nil {
		log.Fatal(err)
	}
	orders := spot.OrderList{*order}
	printer.Print(&orders)
}
