package spot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/UnipayFI/aster-cli/common"
	"github.com/UnipayFI/aster-cli/config"
	"github.com/UnipayFI/aster-cli/exchange"
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
		Run:     orderList,
	}
	orderOpenListCmd = &cobra.Command{
		Use:   "open",
		Short: "List open orders",
		Run:   orderOpenList,
	}
	orderCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create order",
		Long: `Create a new order.
* Supports parameters: symbol, side, type, quantity, quoteOrderQty, timeInForce, price, newClientOrderId, stopPrice`,
		Run: createOrder,
	}
	orderCancelCmd = &cobra.Command{
		Use:   "cancel",
		Short: "Cancel order",
		Long: `Cancel order.
If either orderId or origClientOrderId is provided, the specified order will be canceled.
If only the symbol is passed, all open orders for that trading pair will be canceled.`,
		Run: cancelOrder,
	}
	orderGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Query a single order",
		Long:  `Query a single order by orderId or origClientOrderId.`,
		Run:   getOrder,
	}
)

func InitOrderCmds() []*cobra.Command {
	orderCmd.PersistentFlags().StringP("symbol", "s", "", "symbol")

	orderListCmd.Flags().Int64P("orderId", "i", 0, "orderId")
	orderListCmd.Flags().IntP("limit", "l", 500, "limit, max 1000")
	orderListCmd.Flags().Int64P("startTime", "a", 0, "start time (unix milliseconds)")
	orderListCmd.Flags().Int64P("endTime", "e", 0, "end time (unix milliseconds)")
	orderListCmd.MarkFlagRequired("symbol")

	var side, orderType string
	orderCreateCmd.Flags().StringVarP(&side, "side", "S", "", "BUY or SELL")
	orderCreateCmd.Flags().StringVarP(&orderType, "type", "t", "", "LIMIT, MARKET, STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, TAKE_PROFIT_LIMIT, LIMIT_MAKER")
	orderCreateCmd.Flags().Float64P("quantity", "q", 0, "order quantity")
	orderCreateCmd.Flags().Float64("quoteOrderQty", 0, "quote order quantity (for MARKET orders)")
	orderCreateCmd.Flags().Float64P("price", "p", 0, "order price (required for LIMIT orders)")
	orderCreateCmd.Flags().StringP("timeInForce", "T", "", "GTC, IOC, FOK (default GTC for LIMIT orders)")
	orderCreateCmd.Flags().Float64("stopPrice", 0, "stop price for STOP_LOSS/TAKE_PROFIT orders")
	orderCreateCmd.Flags().String("newClientOrderId", "", "custom order id")
	orderCreateCmd.FParseErrWhitelist = cobra.FParseErrWhitelist{
		UnknownFlags: true,
	}
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
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	limit, _ := cmd.Flags().GetInt("limit")
	startTimeMs, _ := cmd.Flags().GetInt64("startTime")
	endTimeMs, _ := cmd.Flags().GetInt64("endTime")
	orderID, _ := cmd.Flags().GetInt64("orderId")

	var startTime, endTime time.Time
	if startTimeMs != 0 {
		startTime = time.UnixMilli(startTimeMs)
	}
	if endTimeMs != 0 {
		endTime = time.UnixMilli(endTimeMs)
	}

	orders, err := client.GetOrderList(symbol, orderID, startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(orders)
}

func orderOpenList(cmd *cobra.Command, _ []string) {
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orders, err := client.GetOpenOrders(symbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(orders)
}

func createOrder(cmd *cobra.Command, _ []string) {
	_, args, _ := cmd.Root().Find(os.Args[1:])
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	order, err := client.CreateOrder(common.ParseArgs(args))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("order created, orderID:", order.OrderId)
	}
}

func cancelOrder(cmd *cobra.Command, _ []string) {
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
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
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orderID, _ := cmd.Flags().GetInt64("orderId")
	clientOrderID, _ := cmd.Flags().GetString("origClientOrderId")

	order, err := client.GetOrder(symbol, orderID, clientOrderID)
	if err != nil {
		log.Fatal(err)
	}
	orders := spot.OrderList{*order}
	printer.PrintTable(&orders)
}
