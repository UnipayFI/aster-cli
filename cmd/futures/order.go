package futures

import (
	"log"
	"strconv"

	"github.com/UnipayFI/aster-cli/exchange/futures"
	"github.com/UnipayFI/aster-cli/printer"
	asterfutures "github.com/UnipayFI/go-aster/v3/futures"
	"github.com/spf13/cobra"
)

var (
	orderCmd = &cobra.Command{
		Use:   "order",
		Short: "Order management commands",
		Long:  `Manage orders: create, cancel, list, query, force orders, trades.`,
	}

	orderListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all account orders",
		Long: `Get all account orders; active, canceled, or filled.
- Orders not found: status is 'CANCELED' or 'EXPIRED' with no fills and created > 3 days ago
- Orders older than 90 days are not returned

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#all-orders-user_data`,
		Run: orderList,
	}

	orderOpenListCmd = &cobra.Command{
		Use:   "open",
		Short: "List open orders",
		Long: `Get all open orders on a symbol.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#current-all-open-orders-user_data`,
		Run: orderOpenList,
	}

	orderForceCloseCmd = &cobra.Command{
		Use:   "force",
		Short: "Query force orders (liquidation)",
		Long: `Query user's force orders (liquidation orders).
- If "autoCloseType" is not sent, orders with both types will be returned
- If "startTime" is not sent, data within 7 days before "endTime" can be queried

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#users-force-orders-user_data`,
		Run: forceCloseOrder,
	}

	orderCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create a new order",
		Long: `Create a new order.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#new-order-trade`,
		Run: createOrder,
	}

	orderCancelCmd = &cobra.Command{
		Use:   "cancel",
		Short: "Cancel order(s)",
		Long: `Cancel order(s).
If either orderId or orgClientOrderId is provided, the specified order will be canceled.
If only the symbol is passed, all open orders for that trading pair will be canceled.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#cancel-order-trade`,
		Run: cancelOrder,
	}

	orderGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Query a single order",
		Long: `Query a single order by orderId or origClientOrderId.

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#query-order-user_data`,
		Run: getOrder,
	}

	orderTradeCmd = &cobra.Command{
		Use:     "trade",
		Aliases: []string{"trades"},
		Short:   "Query trade history",
		Long: `Get trades for a specific account and symbol.
- If 'startTime' and 'endTime' are both not sent, then the last 7 days' data will be returned
- The time between 'startTime' and 'endTime' cannot be longer than 7 days
- Only support querying trades in the past 6 months

Docs Link: https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#account-trade-list-user_data`,
		Run: orderTrades,
	}
)

func InitOrderCmds() []*cobra.Command {
	orderCmd.PersistentFlags().StringP("symbol", "s", "", "Trading pair symbol")

	orderListCmd.Flags().Int64P("orderId", "i", 0, "Order ID")
	orderListCmd.Flags().IntP("limit", "l", 500, "Number of results (default 500, max 1000)")
	orderListCmd.Flags().StringP("startTime", "a", "", "Start time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	orderListCmd.Flags().StringP("endTime", "e", "", "End time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	orderListCmd.MarkFlagRequired("symbol")

	orderForceCloseCmd.Flags().StringP("autoCloseType", "t", "", "Auto close type: LIQUIDATION or ADL")
	orderForceCloseCmd.Flags().StringP("startTime", "a", "", "Start time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	orderForceCloseCmd.Flags().StringP("endTime", "e", "", "End time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	orderForceCloseCmd.Flags().IntP("limit", "l", 50, "Number of results (default 50, max 100)")

	var side, orderType string
	orderCreateCmd.Flags().StringVarP(&side, "side", "S", "", "BUY or SELL")
	orderCreateCmd.Flags().StringVarP(&orderType, "type", "t", "", "LIMIT, MARKET, STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET")
	orderCreateCmd.Flags().StringP("positionSide", "P", "", "LONG or SHORT (default BOTH for One-way Mode)")
	orderCreateCmd.Flags().StringP("quantity", "q", "", "Order quantity (decimal string)")
	orderCreateCmd.Flags().StringP("price", "p", "", "Order price, required for LIMIT orders (decimal string)")
	orderCreateCmd.Flags().StringP("timeInForce", "T", "", "GTC, IOC, FOK, GTX (default GTC for LIMIT orders)")
	orderCreateCmd.Flags().Bool("reduceOnly", false, "Reduce only order")
	orderCreateCmd.Flags().String("stopPrice", "", "Stop price for STOP/TAKE_PROFIT orders (decimal string)")
	orderCreateCmd.Flags().Bool("closePosition", false, "Close all position")
	orderCreateCmd.Flags().String("activationPrice", "", "Activation price for TRAILING_STOP_MARKET (decimal string)")
	orderCreateCmd.Flags().String("callbackRate", "", "Callback rate for TRAILING_STOP_MARKET, min 0.1 max 5 (decimal string)")
	orderCreateCmd.Flags().String("workingType", "", "MARK_PRICE or CONTRACT_PRICE")
	orderCreateCmd.Flags().Bool("priceProtect", false, "Price protection")
	orderCreateCmd.Flags().String("newClientOrderId", "", "Custom order ID")
	orderCreateCmd.Flags().String("newOrderRespType", "", "ACK, RESULT (default ACK)")
	orderCreateCmd.MarkFlagRequired("symbol")

	orderCancelCmd.Flags().Int64P("orderId", "i", 0, "Order ID")
	orderCancelCmd.Flags().StringP("origClientOrderId", "c", "", "Client order ID")
	orderCancelCmd.MarkFlagRequired("symbol")

	orderGetCmd.Flags().Int64P("orderId", "i", 0, "Order ID")
	orderGetCmd.Flags().StringP("origClientOrderId", "c", "", "Client order ID")
	orderGetCmd.MarkFlagRequired("symbol")

	orderTradeCmd.Flags().StringP("startTime", "a", "", "Start time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	orderTradeCmd.Flags().StringP("endTime", "e", "", "End time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	orderTradeCmd.Flags().Int64P("fromId", "f", 0, "Trade ID to fetch from")
	orderTradeCmd.Flags().IntP("limit", "l", 500, "Number of results (default 500, max 1000)")
	orderTradeCmd.MarkFlagRequired("symbol")

	orderCmd.AddCommand(
		orderListCmd,
		orderOpenListCmd,
		orderForceCloseCmd,
		orderCreateCmd,
		orderCancelCmd,
		orderGetCmd,
		orderTradeCmd,
	)
	return []*cobra.Command{orderCmd}
}

func orderList(cmd *cobra.Command, _ []string) {
	client := newClient()
	symbol, _ := cmd.Flags().GetString("symbol")
	limit, _ := cmd.Flags().GetInt("limit")
	startTimeRaw, _ := cmd.Flags().GetString("startTime")
	endTimeRaw, _ := cmd.Flags().GetString("endTime")
	orderID, _ := cmd.Flags().GetInt64("orderId")
	startTime, err := parseTimeFlag("--startTime", startTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	endTime, err := parseTimeFlag("--endTime", endTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	orders, err := client.GetOrderList(symbol, limit, startTime, endTime, orderID)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(&orders)
}

func orderOpenList(cmd *cobra.Command, _ []string) {
	client := newClient()
	symbol, _ := cmd.Flags().GetString("symbol")
	orders, err := client.GetOpenOrders(symbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(&orders)
}

func forceCloseOrder(cmd *cobra.Command, _ []string) {
	client := newClient()
	symbol, _ := cmd.Flags().GetString("symbol")
	autoCloseType, _ := cmd.Flags().GetString("autoCloseType")
	startTimeRaw, _ := cmd.Flags().GetString("startTime")
	endTimeRaw, _ := cmd.Flags().GetString("endTime")
	limit, _ := cmd.Flags().GetInt("limit")
	startTime, err := parseTimeFlag("--startTime", startTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	endTime, err := parseTimeFlag("--endTime", endTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	orders, err := client.GetForceOrders(symbol, asterfutures.AutoCloseType(autoCloseType), startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(&orders)
}

func createOrder(cmd *cobra.Command, _ []string) {
	client := newClient()
	order, err := client.CreateOrder(buildOrderParams(cmd))
	if err != nil {
		log.Fatal(err)
	}
	orders := futures.OrderList{*order}
	printer.Print(&orders)
}

// buildOrderParams reads the cobra-parsed flags and returns the long-name
// keyed map that exchange/futures.Client.CreateOrder expects. Bool flags are
// encoded as "true"/"false" because the underlying SDK wrapper checks the
// string value.
func buildOrderParams(cmd *cobra.Command) map[string]string {
	params := map[string]string{}
	stringFlags := []string{
		"symbol", "side", "type", "positionSide", "quantity",
		"price", "timeInForce", "stopPrice", "activationPrice",
		"callbackRate", "workingType", "newClientOrderId", "newOrderRespType",
	}
	for _, name := range stringFlags {
		if v, _ := cmd.Flags().GetString(name); v != "" {
			params[name] = v
		}
	}
	boolFlags := []string{"reduceOnly", "closePosition", "priceProtect"}
	for _, name := range boolFlags {
		if cmd.Flags().Changed(name) {
			b, _ := cmd.Flags().GetBool(name)
			params[name] = strconv.FormatBool(b)
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
	orders := futures.OrderList{*order}
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
	orders := futures.OrderList{*order}
	printer.Print(&orders)
}

func orderTrades(cmd *cobra.Command, _ []string) {
	client := newClient()
	symbol, _ := cmd.Flags().GetString("symbol")
	startTimeRaw, _ := cmd.Flags().GetString("startTime")
	endTimeRaw, _ := cmd.Flags().GetString("endTime")
	fromId, _ := cmd.Flags().GetInt64("fromId")
	limit, _ := cmd.Flags().GetInt("limit")
	startTime, err := parseTimeFlag("--startTime", startTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	endTime, err := parseTimeFlag("--endTime", endTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	trades, err := client.GetTrades(symbol, startTime, endTime, fromId, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(&trades)
}
