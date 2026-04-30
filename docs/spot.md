# Spot Module

> Every command in this module accepts a global `--json` flag that prints the
> raw API response as indented JSON instead of a table. See the [JSON output](#json-output) section.

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
- [JSON output](#json-output)

## Account - Show account info
Docs Link: <https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#account-information-user_data>

Exec: `./aster-cli spot account`
```shell
┌──────────┬───────────┬──────────────┬─────────────┬─────────────────────┐
│ FEE TIER │ CAN TRADE │ CAN WITHDRAW │ CAN DEPOSIT │     UPDATE TIME     │
├──────────┼───────────┼──────────────┼─────────────┼─────────────────────┤
│ 0        │ true      │ true         │ true        │ 2026-04-30 12:00:00 │
└──────────┴───────────┴──────────────┴─────────────┴─────────────────────┘
```

## Balance - Show account balances
Docs Link: <https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#account-information-user_data>

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
Docs Link: <https://asterdex.github.io/aster-api-website/spot-v3/market-data/#get-symbol-fees>

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
Docs Link: <https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#place-order-trade>

Create a new order. The created order is printed back as a single-row table
(or JSON with `--json`).

**Market Order:**
```shell
./aster-cli spot order create --symbol=ETHUSDT --side=SELL --type=MARKET --quantity=0.003
```

**Limit Order:**
```shell
./aster-cli spot order create --symbol=ETHUSDT --side=BUY --type=LIMIT --quantity=0.003 --price=2500
```

**Supported parameters:**
- `--symbol, -s`: Trading pair symbol (required)
- `--side, -S`: BUY or SELL
- `--type, -t`: LIMIT, MARKET, STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET
- `--quantity, -q`: Order quantity (decimal string)
- `--quoteOrderQty`: Quote order quantity (decimal string, MARKET orders)
- `--price, -p`: Order price (decimal string, required for LIMIT orders)
- `--timeInForce, -T`: GTC, IOC, FOK (default GTC for LIMIT orders)
- `--stopPrice`: Stop price (decimal string) for STOP/TAKE_PROFIT orders
- `--newClientOrderId`: Custom order ID

### Order - List open orders
Docs Link: <https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#current-open-orders-user_data>

Exec: `./aster-cli spot order open`

Or for a specific symbol:
```shell
./aster-cli spot order open --symbol=ETHUSDT
```
```shell
┌───────────┬────────────────────────┬─────────┬──────┬───────┬────────┬───────┬───────────┬──────────┬──────────────┬───────────┬─────┬─────────────────────┬─────────────────────┐
│ ORDER ID  │    CLIENT ORDER ID     │ SYMBOL  │ SIDE │ TYPE  │ STATUS │ PRICE │ AVG PRICE │ QUANTITY │ EXECUTED QTY │ CUM QUOTE │ TIF │        TIME         │     UPDATE TIME     │
├───────────┼────────────────────────┼─────────┼──────┼───────┼────────┼───────┼───────────┼──────────┼──────────────┼───────────┼─────┼─────────────────────┼─────────────────────┤
│ 168166704 │ 8dEU6uBAZoJzR4SD8fCt3j │ ETHUSDT │ BUY  │ LIMIT │ NEW    │ 2500  │ 0         │ 0.003    │ 0            │ 0         │ GTC │ 2026-04-30 03:57:55 │ 2026-04-30 03:57:55 │
└───────────┴────────────────────────┴─────────┴──────┴───────┴────────┴───────┴───────────┴──────────┴──────────────┴───────────┴─────┴─────────────────────┴─────────────────────┘
```

### Order - List orders
Docs Link: <https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#query-all-orders-user_data>

Exec: `./aster-cli spot order list --symbol=ETHUSDT`

**Supported parameters:**
- `--symbol, -s`: Trading pair symbol (required)
- `--orderId, -i`: Start from this orderId (returns orders >= orderId)
- `--limit, -l`: Number of results (default 500, max 1000)
- `--startTime, -a`: Start time (unix ms or "YYYY-MM-DD HH:MM:SS")
- `--endTime, -e`: End time (unix ms or "YYYY-MM-DD HH:MM:SS")

If the value contains spaces, wrap it in quotes, e.g. `--startTime "2025-12-18 04:16:21"`. Date time strings are parsed in local timezone.

### Order - Get Order
Docs Link: <https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#query-order-user_data>

Query a single order by orderId or origClientOrderId.
```shell
./aster-cli spot order get --symbol=ETHUSDT --orderId=13557621683
```
Or:
```shell
./aster-cli spot order get --symbol=ETHUSDT --origClientOrderId=xxxxx
```

### Order - Cancel order by ID
Docs Link: <https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#cancel-order-trade>

Exec: `./aster-cli spot order cancel --symbol=ETHUSDT --orderId=xxxxx`

Or by client order ID:
```shell
./aster-cli spot order cancel --symbol=ETHUSDT --origClientOrderId=xxxxx
```
The canceled order is printed back as a single-row table (or JSON with `--json`).

### Order - Cancel all orders by symbol
Docs Link: <https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#cancel-order-trade>

Exec: `./aster-cli spot order cancel --symbol=ETHUSDT`

If neither `--orderId` nor `--origClientOrderId` is provided, all open orders for the symbol will be canceled. The response is the API's `{code, msg}` envelope.

## Trade - Query user trades
Docs Link: <https://asterdex.github.io/aster-api-website/spot-v3/account%26trades/#account-trade-history-user_data>

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
│ 822714 │ 163314308 │ ETHUSDT │ BUY  │ 2960.1  │ 0.003    │ 8.88           │ 0.0000012 ETH   │ 2026-04-29 11:30:18 │ false │
│ 879877 │ 168165256 │ ETHUSDT │ SELL │ 2831.74 │ 0.003    │ 8.49           │ 0.00339809 USDT │ 2026-04-30 03:56:56 │ false │
└────────┴───────────┴─────────┴──────┴─────────┴──────────┴────────────────┴─────────────────┴─────────────────────┴───────┘
```

## JSON output

Add `--json` (a global flag, position doesn't matter) to receive the raw V3
API response as indented JSON. Useful for piping into `jq`.

```shell
./aster-cli spot account --json
```
```json
{
  "feeTier": 0,
  "canTrade": true,
  "canDeposit": true,
  "canWithdraw": true,
  "canBurnAsset": false,
  "updateTime": 1746000000000,
  "balances": [
    {
      "asset": "USDT",
      "free": "2.1197",
      "locked": "0"
    }
  ]
}
```
