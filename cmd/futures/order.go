package futures

import (
	"fmt"
	"log"
	"os"

	"github.com/UnipayFI/aster-cli/common"
	"github.com/UnipayFI/aster-cli/config"
	"github.com/UnipayFI/aster-cli/exchange"
	"github.com/UnipayFI/aster-cli/exchange/futures"
	"github.com/UnipayFI/aster-cli/printer"
	asterfutures "github.com/UnipayFI/go-aster/futures"
	"github.com/spf13/cobra"
)

var (
	orderCmd = &cobra.Command{
		Use:   "order",
		Short: "Order management commands",
		Long:  `Manage orders: create, cancel, list, query, force orders, trades.`,
	}

	// order list
	orderListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all account orders",
		Long: `Get all account orders; active, canceled, or filled.
- Orders not found: status is 'CANCELED' or 'EXPIRED' with no fills and created > 3 days ago
- Orders older than 90 days are not returned`,
		Run: orderList,
	}

	// order open
	orderOpenListCmd = &cobra.Command{
		Use:   "open",
		Short: "List open orders",
		Long:  `Get all open orders on a symbol.`,
		Run:   orderOpenList,
	}

	// order force
	orderForceCloseCmd = &cobra.Command{
		Use:   "force",
		Short: "Query force orders (liquidation)",
		Long: `Query user's force orders (liquidation orders).
- If "autoCloseType" is not sent, orders with both types will be returned
- If "startTime" is not sent, data within 7 days before "endTime" can be queried`,
		Run: forceCloseOrder,
	}

	// order create
	orderCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create a new order",
		Long:    `Create a new order.`,
		Run:     createOrder,
	}

	// order cancel
	orderCancelCmd = &cobra.Command{
		Use:   "cancel",
		Short: "Cancel order(s)",
		Long: `Cancel order(s).
If either orderId or orgClientOrderId is provided, the specified order will be canceled.
If only the symbol is passed, all open orders for that trading pair will be canceled.`,
		Run: cancelOrder,
	}

	// order get
	orderGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Query a single order",
		Long:  `Query a single order by orderId or origClientOrderId.`,
		Run:   getOrder,
	}

	// order trade
	orderTradeCmd = &cobra.Command{
		Use:     "trade",
		Aliases: []string{"trades"},
		Short:   "Query trade history",
		Long: `Get trades for a specific account and symbol.
- If 'startTime' and 'endTime' are both not sent, then the last 7 days' data will be returned
- The time between 'startTime' and 'endTime' cannot be longer than 7 days
- Only support querying trades in the past 6 months`,
		Run: orderTrades,
	}
)

func InitOrderCmds() []*cobra.Command {
	orderCmd.PersistentFlags().StringP("symbol", "s", "", "Trading pair symbol")

	// order list flags
	orderListCmd.Flags().Int64P("orderId", "i", 0, "Order ID")
	orderListCmd.Flags().IntP("limit", "l", 500, "Number of results (default 500, max 1000)")
	orderListCmd.Flags().StringP("startTime", "a", "", "Start time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	orderListCmd.Flags().StringP("endTime", "e", "", "End time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	orderListCmd.MarkFlagRequired("symbol")

	// order force flags
	orderForceCloseCmd.Flags().StringP("autoCloseType", "t", "", "Auto close type: LIQUIDATION or ADL")
	orderForceCloseCmd.Flags().StringP("startTime", "a", "", "Start time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	orderForceCloseCmd.Flags().StringP("endTime", "e", "", "End time (unix ms or \"YYYY-MM-DD HH:MM:SS\")")
	orderForceCloseCmd.Flags().IntP("limit", "l", 50, "Number of results (default 50, max 100)")

	// order create flags
	var side, orderType string
	orderCreateCmd.Flags().StringVarP(&side, "side", "S", "", "BUY or SELL")
	orderCreateCmd.Flags().StringVarP(&orderType, "type", "t", "", "LIMIT, MARKET, STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET")
	orderCreateCmd.Flags().StringP("positionSide", "P", "", "LONG or SHORT (default BOTH for One-way Mode)")
	orderCreateCmd.Flags().Float64P("quantity", "q", 0, "Order quantity")
	orderCreateCmd.Flags().Float64P("price", "p", 0, "Order price (required for LIMIT orders)")
	orderCreateCmd.Flags().StringP("timeInForce", "T", "", "GTC, IOC, FOK, GTX (default GTC for LIMIT orders)")
	orderCreateCmd.Flags().Bool("reduceOnly", false, "Reduce only order")
	orderCreateCmd.Flags().Float64("stopPrice", 0, "Stop price for STOP/TAKE_PROFIT orders")
	orderCreateCmd.Flags().Bool("closePosition", false, "Close all position")
	orderCreateCmd.Flags().Float64("activationPrice", 0, "Activation price for TRAILING_STOP_MARKET")
	orderCreateCmd.Flags().Float64("callbackRate", 0, "Callback rate for TRAILING_STOP_MARKET (min 0.1, max 5)")
	orderCreateCmd.Flags().String("workingType", "", "MARK_PRICE or CONTRACT_PRICE")
	orderCreateCmd.Flags().Bool("priceProtect", false, "Price protection")
	orderCreateCmd.Flags().String("newClientOrderId", "", "Custom order ID")
	orderCreateCmd.Flags().String("newOrderRespType", "", "ACK, RESULT (default ACK)")
	orderCreateCmd.FParseErrWhitelist = cobra.FParseErrWhitelist{
		UnknownFlags: true,
	}
	orderCreateCmd.MarkFlagRequired("symbol")

	// order cancel flags
	orderCancelCmd.Flags().Int64P("orderId", "i", 0, "Order ID")
	orderCancelCmd.Flags().StringP("origClientOrderId", "c", "", "Client order ID")
	orderCancelCmd.MarkFlagRequired("symbol")

	// order get flags
	orderGetCmd.Flags().Int64P("orderId", "i", 0, "Order ID")
	orderGetCmd.Flags().StringP("origClientOrderId", "c", "", "Client order ID")
	orderGetCmd.MarkFlagRequired("symbol")

	// order trade flags
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
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	limit, _ := cmd.Flags().GetInt("limit")
	startTimeRaw, _ := cmd.Flags().GetString("startTime")
	endTimeRaw, _ := cmd.Flags().GetString("endTime")
	orderID, _ := cmd.Flags().GetInt64("orderId")
	startTime, _, err := common.ParseTimeFlagUnixMilli("--startTime", startTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	endTime, _, err := common.ParseTimeFlagUnixMilli("--endTime", endTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	orders, err := client.GetOrderList(symbol, limit, startTime, endTime, orderID)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func orderOpenList(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orders, err := client.GetOpenOrders(symbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func forceCloseOrder(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	autoCloseType, _ := cmd.Flags().GetString("autoCloseType")
	startTimeRaw, _ := cmd.Flags().GetString("startTime")
	endTimeRaw, _ := cmd.Flags().GetString("endTime")
	limit, _ := cmd.Flags().GetInt("limit")
	startTime, _, err := common.ParseTimeFlagUnixMilli("--startTime", startTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	endTime, _, err := common.ParseTimeFlagUnixMilli("--endTime", endTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	orders, err := client.GetForceOrders(symbol, asterfutures.AutoCloseType(autoCloseType), startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func createOrder(cmd *cobra.Command, _ []string) {
	_, args, _ := cmd.Root().Find(os.Args[1:])
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	order, err := client.CreateOrder(common.ParseArgs(args))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("order created, orderID:", order.OrderId)
	}
}

func cancelOrder(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orderID, _ := cmd.Flags().GetInt64("orderId")
	clientOrderID, _ := cmd.Flags().GetString("origClientOrderId")

	var err error
	if orderID == 0 && clientOrderID == "" {
		err = client.CancelAllOrders(symbol)
	} else {
		err = client.CancelOrder(symbol, orderID, clientOrderID)
	}
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("order canceled")
	}
}

func getOrder(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orderID, _ := cmd.Flags().GetInt64("orderId")
	clientOrderID, _ := cmd.Flags().GetString("origClientOrderId")

	order, err := client.GetOrder(symbol, orderID, clientOrderID)
	if err != nil {
		log.Fatal(err)
	}
	// Wrap single order in OrderList for table printing
	orders := futures.OrderList{*order}
	printer.PrintTable(&orders)
}

func orderTrades(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	startTimeRaw, _ := cmd.Flags().GetString("startTime")
	endTimeRaw, _ := cmd.Flags().GetString("endTime")
	fromId, _ := cmd.Flags().GetInt64("fromId")
	limit, _ := cmd.Flags().GetInt("limit")
	startTime, _, err := common.ParseTimeFlagUnixMilli("--startTime", startTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	endTime, _, err := common.ParseTimeFlagUnixMilli("--endTime", endTimeRaw)
	if err != nil {
		log.Fatal(err)
	}
	trades, err := client.GetTrades(symbol, startTime, endTime, fromId, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&trades)
}
