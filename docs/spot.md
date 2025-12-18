# Spot Module

## Quick Navigation
- [Account](#account---show-account-info)
- [Balance](#balance---show-account-balances)
- [Commission-rate](#commission-rate---get-commission-rate)
- [Order](#order)
  - [Create Order](#order---create-order)
  - [List Open Orders](#order---list-open-orders)
  - [List Orders](#order---list-orders)
  - [Get Order](#order---get-order)
  - [Cancel Order by ID](#order---cancel-order-by-id)
  - [Cancel All Orders by Symbol](#order---cancel-all-orders-by-symbol)
- [Trade](#trade---query-user-trades)

## Account - Show account info
Exec: `./aster-cli spot account`
```shell
┌──────────┬───────────┬──────────────┬─────────────┬─────────────────────┐
│ FEE TIER │ CAN TRADE │ CAN WITHDRAW │ CAN DEPOSIT │     UPDATE TIME     │
├──────────┼───────────┼──────────────┼─────────────┼─────────────────────┤
│ 0        │ true      │ true         │ true        │ 2025-12-13 13:32:40 │
└──────────┴───────────┴──────────────┴─────────────┴─────────────────────┘
```

## Balance - Show account balances
Exec: `./aster-cli spot balance`

Shows only non-zero balances.
```shell
┌───────┬────────────┬────────┐
│ ASSET │    FREE    │ LOCKED │
├───────┼────────────┼────────┤
│ USDT  │ 2.1197     │ 0      │
│ ETH   │ 0.00309752 │ 0      │
└───────┴────────────┴────────┘
```

## Commission-rate - Get commission rate
Exec: `./aster-cli spot commission-rate --symbol=BTCUSDT`
```shell
┌─────────┬───────────────────────┬───────────────────────┐
│ SYMBOL  │ MAKER COMMISSION RATE │ TAKER COMMISSION RATE │
├─────────┼───────────────────────┼───────────────────────┤
│ BTCUSDT │ 0.00005               │ 0.0004                │
└─────────┴───────────────────────┴───────────────────────┘
```

## Order

### Order - Create Order
Create a new order. Supports various order types.

**Market Order:**

Exec: `./aster-cli spot order create --symbol=ETHUSDT --side=SELL --type=MARKET --quantity=0.003`

```shell
order created, orderID: 168165256
```

**Limit Order:**

Exec: `./aster-cli spot order create --symbol=ETHUSDT --side=SELL --type=MARKET --quantity=0.003`

```shell
order created, orderID: 168166704
```

**Supported parameters:**
- `--symbol, -s`: Trading pair symbol (required)
- `--side, -S`: BUY or SELL
- `--type, -t`: LIMIT, MARKET, STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, TAKE_PROFIT_LIMIT, LIMIT_MAKER
- `--quantity, -q`: Order quantity
- `--quoteOrderQty`: Quote order quantity (for MARKET orders)
- `--price, -p`: Order price (required for LIMIT orders)
- `--timeInForce, -T`: GTC, IOC, FOK
- `--stopPrice`: Stop price for STOP_LOSS/TAKE_PROFIT orders
- `--newClientOrderId`: Custom order ID

### Order - List open orders
Exec: `./aster-cli spot order open`

Or for a specific symbol:
```shell
./aster-cli spot order open --symbol=ETHUSDT
```
```shell
┌───────────┬────────────────────────┬─────────┬──────┬───────┬────────┬───────┬───────────┬──────────┬──────────────┬───────────┬─────┬─────────────────────┬─────────────────────┐
│ ORDER ID  │    CLIENT ORDER ID     │ SYMBOL  │ SIDE │ TYPE  │ STATUS │ PRICE │ AVG PRICE │ QUANTITY │ EXECUTED QTY │ CUM QUOTE │ TIF │        TIME         │     UPDATE TIME     │
├───────────┼────────────────────────┼─────────┼──────┼───────┼────────┼───────┼───────────┼──────────┼──────────────┼───────────┼─────┼─────────────────────┼─────────────────────┤
│ 168166704 │ 8dEU6uBAZoJzR4SD8fCt3j │ ETHUSDT │ BUY  │ LIMIT │ NEW    │ 2500  │ 0         │ 0.003    │ 0            │ 0         │ GTC │ 2025-12-18 03:57:55 │ 2025-12-18 03:57:55 │
└───────────┴────────────────────────┴─────────┴──────┴───────┴────────┴───────┴───────────┴──────────┴──────────────┴───────────┴─────┴─────────────────────┴─────────────────────┘
```

### Order - List orders
Exec: `./aster-cli spot order list --symbol=ETHUSDT`

**Supported parameters:**
- `--symbol, -s`: Trading pair symbol (required)
- `--orderId, -i`: Start from this orderId (returns orders >= orderId)
- `--limit, -l`: Number of results (default 500, max 1000)
- `--startTime, -a`: Start time (unix milliseconds)
- `--endTime, -e`: End time (unix milliseconds)

```shell
┌───────────┬────────────────────────┬─────────┬──────┬────────┬──────────┬─────────┬───────────┬──────────┬──────────────┬───────────┬─────┬─────────────────────┬─────────────────────┐
│ ORDER ID  │    CLIENT ORDER ID     │ SYMBOL  │ SIDE │  TYPE  │  STATUS  │  PRICE  │ AVG PRICE │ QUANTITY │ EXECUTED QTY │ CUM QUOTE │ TIF │        TIME         │     UPDATE TIME     │
├───────────┼────────────────────────┼─────────┼──────┼────────┼──────────┼─────────┼───────────┼──────────┼──────────────┼───────────┼─────┼─────────────────────┼─────────────────────┤
│ 158402974 │ QiKytJmimeCEhTxlUltoi4 │ ETHUSDT │ BUY  │ LIMIT  │ CANCELED │ 3000    │ 0         │ 0.003    │ 0            │ 0         │ GTC │ 2025-12-14 13:35:11 │ 2025-12-14 13:35:11 │
│ 163072581 │ Y413NjjqgklLYuBO8srYR6 │ ETHUSDT │ BUY  │ LIMIT  │ CANCELED │ 2500    │ 0         │ 0.003    │ 0            │ 0         │ GTC │ 2025-12-16 09:06:28 │ 2025-12-16 09:11:17 │
│ 163314308 │ Z6SUTAKChfMmpJVfuhtV1i │ ETHUSDT │ BUY  │ LIMIT  │ FILLED   │ 3000    │ 2960.1    │ 0.003    │ 0.003        │ 8.8803    │ GTC │ 2025-12-16 11:30:18 │ 2025-12-16 11:30:18 │
│ 168165256 │ i54prsxLjATWImm86ls8Dz │ ETHUSDT │ SELL │ MARKET │ FILLED   │ 2831.74 │ 2831.74   │ 0.003    │ 0.003        │ 8.49522   │ GTC │ 2025-12-18 03:56:56 │ 2025-12-18 03:56:56 │
│ 168166704 │ 8dEU6uBAZoJzR4SD8fCt3j │ ETHUSDT │ BUY  │ LIMIT  │ NEW      │ 2500    │ 0         │ 0.003    │ 0            │ 0         │ GTC │ 2025-12-18 03:57:55 │ 2025-12-18 03:57:55 │
└───────────┴────────────────────────┴─────────┴──────┴─
```

### Order - Get Order
Query a single order by orderId or origClientOrderId.
```shell
./aster-cli spot order get --symbol=ETHUSDT --orderId=13557621683
```
Or:
```shell
./aster-cli spot order get --symbol=ETHUSDT --origClientOrderId=xxxxx
```

### Order - Cancel order by ID
Exec: `./aster-cli spot order cancel --symbol=ETHUSDT --orderId=xxxxx`

Or by client order ID:
```shell
./aster-cli spot order cancel --symbol=ETHUSDT --origClientOrderId=xxxxx
```

### Order - Cancel all orders by symbol
Exec: `./aster-cli spot order cancel --symbol=ETHUSDT`

If neither `--orderId` nor `--origClientOrderId` is provided, all open orders for the symbol will be canceled.

## Trade - Query user trades
Exec: `./aster-cli spot trade list --symbol=ETHUSDT`

**Supported parameters:**

- `--symbol, -s`: Trading pair symbol (required)
- `--orderId, -o`: Order ID to filter trades
- `--fromId, -f`: Trade ID to fetch from
- `--limit, -l`: Number of results (default 500, max 1000)

```shell
┌────────┬───────────┬─────────┬──────┬─────────┬──────────┬────────────────┬─────────────────┬─────────────────────┬───────┐
│   ID   │ ORDER ID  │ SYMBOL  │ SIDE │  PRICE  │ QUANTITY │ QUOTE QUANTITY │   COMMISSION    │        TIME         │ MAKER │
├────────┼───────────┼─────────┼──────┼─────────┼──────────┼────────────────┼─────────────────┼─────────────────────┼───────┤
│ 736272 │ 156389826 │ ETHUSDT │ BUY  │ 3100.75 │ 0.0032   │ 9.92           │ 0.00000128 ETH  │ 2025-12-13 14:32:02 │ false │
│ 750758 │ 158004192 │ ETHUSDT │ SELL │ 3107.31 │ 0.0031   │ 9.63           │ 0.00385306 USDT │ 2025-12-14 09:31:27 │ false │
│ 822714 │ 163314308 │ ETHUSDT │ BUY  │ 2960.1  │ 0.003    │ 8.88           │ 0.0000012 ETH   │ 2025-12-16 11:30:18 │ false │
│ 879877 │ 168165256 │ ETHUSDT │ SELL │ 2831.74 │ 0.003    │ 8.49           │ 0.00339809 USDT │ 2025-12-18 03:56:56 │ false │
└────────┴───────────┴─────────┴──────┴─────────┴──────────┴────────────────┴─────────────────┴─────────────────────┴───────┘
```
