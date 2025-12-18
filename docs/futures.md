# Futures Module

## Quick Navigation
- [Account](#account)
  - [Show account balances](#account---show-account-balances)
  - [Show account info](#account---show-account-info)
  - [Show commission rate](#account---show-commission-rate)
  - [Query income history](#account---query-income-history)
  - [Multi-assets mode](#account---multi-assets-mode)
- [Order](#order)
  - [Create Order](#order---create-order)
  - [List open orders](#order---list-open-orders)
  - [List orders](#order---list-orders)
  - [Get Order](#order---get-order)
  - [Query trades](#order---query-trades)
  - [Query force orders](#order---query-force-orders)
  - [Cancel order by ID](#order---cancel-order-by-id)
  - [Cancel all orders by symbol](#order---cancel-all-orders-by-symbol)
- [Position](#position)
  - [List positions](#position---list-positions)
  - [Show position risk](#position---show-position-risk)
  - [Position mode](#position---position-mode)
  - [Set position margin](#position---set-position-margin)
  - [Margin history](#position---margin-history)
  - [ADL quantile](#position---adl-quantile)
- [Symbol](#symbol)
  - [Set leverage](#symbol---set-leverage)
  - [Set margin type](#symbol---set-margin-type)
  - [Leverage bracket](#symbol---leverage-bracket)
- [Funding](#funding)
  - [Funding info](#funding---funding-info)
  - [Funding rate history](#funding---funding-rate-history)

## Account

### Account - Show account balances
Exec: `./aster-cli futures account balances`
```shell
┌──────────┬───────────┬──────────────────────┬──────────────┬───────────────────┬─────────────────────┐
│  ASSET   │  BALANCE  │ CROSS WALLET BALANCE │ CROSS UN PNL │ AVAILABLE BALANCE │ MAX WITHDRAW AMOUNT │
├──────────┼───────────┼──────────────────────┼──────────────┼───────────────────┼─────────────────────┤
│ CDL      │ 0         │ 0                    │ 0            │ 96.82822961       │ 0                   │
│ FBTC     │ 0         │ 0                    │ 0            │ 0.00006723        │ 0                   │
│ BONUSUSD │ 0         │ 0                    │ 0            │ 6.68838987        │ 0                   │
│ CAKE     │ 0         │ 0                    │ 0            │ 3.09168527        │ 0                   │
│ WBETH    │ 0         │ 0                    │ 0            │ 0.00197382        │ 0                   │
│ BTC      │ 0         │ 0                    │ 0            │ 0.00007349        │ 0                   │
│ TWT      │ 0         │ 0                    │ 0            │ 3.98481436        │ 0                   │
│ SOLVBTC  │ 0         │ 0                    │ 0            │ 0.00007021        │ 0                   │
│ USDCE    │ 0         │ 0                    │ 0            │ 6.68845694        │ 0                   │
│ FORM     │ 0         │ 0                    │ 0            │ 12.20320237       │ 0                   │
│ LISUSD   │ 0         │ 0                    │ 0            │ 6.09135015        │ 0                   │
│ USDBC    │ 0         │ 0                    │ 0            │ 6.68845694        │ 0                   │
│ STONE    │ 0         │ 0                    │ 0            │ 0.00204958        │ 0                   │
│ LISTA    │ 0         │ 0                    │ 0            │ 24.76868462       │ 0                   │
│ RSETH    │ 0         │ 0                    │ 0            │ 0.00202942        │ 0                   │
│ ASTER    │ 0         │ 0                    │ 0            │ 8.19088379        │ 0                   │
│ JLP      │ 0         │ 0                    │ 0            │ 1.3643923         │ 0                   │
│ SOL      │ 0         │ 0                    │ 0            │ 0.0494919         │ 0                   │
│ BNB      │ 0         │ 0                    │ 0            │ 0.00765242        │ 0                   │
│ ETH      │ 0         │ 0                    │ 0            │ 0.00225104        │ 0                   │
│ USDF     │ 0         │ 0                    │ 0            │ 6.6904639         │ 0                   │
│ USDT     │ 8.6922585 │ 8.6922585            │ -0.30366475  │ 6.6904639         │ 6.6904639           │
│ AFEE     │ 0         │ 0                    │ 0            │ 3.34469658        │ 0                   │
│ CUSDT    │ 0         │ 0                    │ 0            │ 284.83739457      │ 0                   │
│ PUMPBTC  │ 0         │ 0                    │ 0            │ 0.00003864        │ 0                   │
│ USD1     │ 0         │ 0                    │ 0            │ 6.5526223         │ 0                   │
│ USDC     │ 0         │ 0                    │ 0            │ 6.68845694        │ 0                   │
│ ASBNB    │ 0         │ 0                    │ 0            │ 0.00720871        │ 0                   │
│ SLISBNB  │ 0         │ 0                    │ 0            │ 0.00706604        │ 0                   │
│ VUSDT    │ 0         │ 0                    │ 0            │ 284.10015961      │ 0                   │
│ BUSD     │ 0         │ 0                    │ 0            │ 6.68905871        │ 0                   │
└──────────┴───────────┴──────────────────────┴──────────────┴───────────────────┴─────────────────────┘
```

### Account - Show account info
Exec: `./aster-cli futures account info`
```shell
┌──────────┬───────────┬─────────────┬──────────────┬──────────────────────┬───────────────────┐
│ FEE TIER │ CAN TRADE │ CAN DEPOSIT │ CAN WITHDRAW │ TOTAL WALLET BALANCE │ AVAILABLE BALANCE │
├──────────┼───────────┼─────────────┼──────────────┼──────────────────────┼───────────────────┤
│ 0        │ true      │ true        │ true         │ 8.688694934782754    │ 6.68772242        │
└──────────┴───────────┴─────────────┴──────────────┴──────────────────────┴───────────────────┘
```

### Account - Show commission rate
Exec: `./aster-cli futures account commission-rate --symbol=BTCUSDT`
```shell
┌─────────┬───────────────────────┬───────────────────────┐
│ SYMBOL  │ MAKER COMMISSION RATE │ TAKER COMMISSION RATE │
├─────────┼───────────────────────┼───────────────────────┤
│ BTCUSDT │ 0.00005               │ 0.0004                │
└─────────┴───────────────────────┴───────────────────────┘
```

### Account - Query income history
Exec: `./aster-cli futures account income`

**Supported parameters:**
- `--symbol, -s`: Trading pair symbol
- `--incomeType, -t`: Income type (e.g., REALIZED_PNL, COMMISSION, FUNDING_FEE)
- `--startTime, -a`: Start time (unix ms or "YYYY-MM-DD HH:MM:SS")
- `--endTime, -e`: End time (unix ms or "YYYY-MM-DD HH:MM:SS")
- `--limit, -l`: Number of results (default 100, max 1000)

If the value contains spaces, wrap it in quotes, e.g. `--startTime "2025-12-18 04:16:21"`. Date time strings are parsed in local timezone.

```shell
┌───────┬─────────────┬─────────────────────────┬─────────────────────────┬─────────┬─────────────────────┬────────────────────┬──────────┐
│ ASSET │   INCOME    │       INCOME TYPE       │          INFO           │ SYMBOL  │        TIME         │      TRAN ID       │ TRADE ID │
├───────┼─────────────┼─────────────────────────┼─────────────────────────┼─────────┼─────────────────────┼────────────────────┼──────────┤
│ USDT  │ -0.00012014 │ FUNDING_FEE             │ FUNDING_FEE             │ ETHUSDT │ 2025-12-18 00:00:01 │ 538239185439545510 │          │
│ USDT  │ -1          │ TRANSFER_FUTURE_TO_SPOT │ TRANSFER_FUTURE_TO_SPOT │         │ 2025-12-18 02:20:34 │ 50548639           │          │
│ USDT  │ -1          │ TRANSFER_FUTURE_TO_SPOT │ TRANSFER_FUTURE_TO_SPOT │         │ 2025-12-18 02:21:23 │ 50548682           │          │
│ USDT  │ -1          │ TRANSFER_FUTURE_TO_SPOT │ TRANSFER_FUTURE_TO_SPOT │         │ 2025-12-18 03:08:26 │ 50552972           │          │
│ USDT  │ 1           │ TRANSFER_SPOT_TO_FUTURE │ TRANSFER_SPOT_TO_FUTURE │         │ 2025-12-18 03:08:26 │ 50552973           │          │
└───────┴─────────────┴─────────────────────────┴─────────────────────────┴─────────┴─────────────────────┴────────────────────┴──────────┘
```

### Account - Multi-assets mode
**Show current mode:**
```shell
./aster-cli futures account multi-assets-mode show
```

**Enable/Disable multi-assets mode:**
```shell
./aster-cli futures account multi-assets-mode set --multiAssetsMargin=true
./aster-cli futures account multi-assets-mode set --multiAssetsMargin=false
```

## Order

### Order - Create Order
Create a new futures order. Supports various order types.

**Market Order (Open Short):**
```shell
./aster-cli futures order create --symbol=ETHUSDT --side=SELL --positionSide=SHORT --type=MARKET --quantity=0.1
```

**Limit Order (Open Long):**
```shell
./aster-cli futures order create --symbol=ETHUSDT --side=BUY --positionSide=LONG --type=LIMIT --timeInForce=GTC --quantity=0.01 --price=4000
```

**Reduce Short Position (Market):**
```shell
./aster-cli futures order create --symbol=ETHUSDT --side=BUY --positionSide=SHORT --type=MARKET --quantity=1.0
```

**Reduce Long Position (Limit):**
```shell
./aster-cli futures order create --symbol=ETHUSDT --side=SELL --positionSide=LONG --type=LIMIT --timeInForce=GTC --price=4000.0 --quantity=0.01
```

**Supported parameters:**
- `--symbol, -s`: Trading pair symbol (required)
- `--side, -S`: BUY or SELL
- `--type, -t`: LIMIT, MARKET, STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET
- `--positionSide, -P`: LONG or SHORT (default BOTH for One-way Mode)
- `--quantity, -q`: Order quantity
- `--price, -p`: Order price (required for LIMIT orders)
- `--timeInForce, -T`: GTC, IOC, FOK, GTX
- `--reduceOnly`: Reduce only order
- `--stopPrice`: Stop price for STOP/TAKE_PROFIT orders
- `--closePosition`: Close all position
- `--activationPrice`: Activation price for TRAILING_STOP_MARKET
- `--callbackRate`: Callback rate for TRAILING_STOP_MARKET (min 0.1, max 5)
- `--workingType`: MARK_PRICE or CONTRACT_PRICE
- `--priceProtect`: Price protection
- `--newClientOrderId`: Custom order ID

### Order - List open orders
Exec: `./aster-cli futures order open`

Or for a specific symbol:
```shell
./aster-cli futures order open --symbol=ETHUSDT
```
```shell
┌─────────────┬─────────┬──────┬───────┬───────────────┬────────┬───────┬───────────┬──────────┬──────────────┬───────────┬─────┬─────────────────────┬─────────────────────┐
│  ORDER ID   │ SYMBOL  │ SIDE │ TYPE  │ POSITION SIDE │ STATUS │ PRICE │ AVG PRICE │ QUANTITY │ EXECUTED QTY │ CUM QUOTE │ TIF │        TIME         │     UPDATE TIME     │
├─────────────┼─────────┼──────┼───────┼───────────────┼────────┼───────┼───────────┼──────────┼──────────────┼───────────┼─────┼─────────────────────┼─────────────────────┤
│ 12929613862 │ ETHUSDT │ SELL │ LIMIT │ BOTH          │ NEW    │ 4000  │ 0         │ 0.003    │ 0            │ 0         │ GTC │ 2025-12-18 05:26:45 │ 2025-12-18 05:26:45 │
└─────────────┴─────────┴──────┴───────┴───────────────┴────────┴───────┴───────────┴──────────┴──────────────┴───────────┴─────┴─────────────────────┴─────────────────────┘
```

### Order - List orders
Exec: `./aster-cli futures order list --symbol=ETHUSDT`

**Supported parameters:**
- `--symbol, -s`: Trading pair symbol (required)
- `--orderId, -i`: Start from this Order ID (returns orders >= orderId)
- `--limit, -l`: Number of results (default 500, max 1000)
- `--startTime, -a`: Start time (unix ms or "YYYY-MM-DD HH:MM:SS")
- `--endTime, -e`: End time (unix ms or "YYYY-MM-DD HH:MM:SS")

```shell
┌─────────────┬─────────┬──────┬────────┬───────────────┬────────┬───────┬───────────┬──────────┬──────────────┬───────────┬─────┬─────────────────────┬─────────────────────┐
│  ORDER ID   │ SYMBOL  │ SIDE │  TYPE  │ POSITION SIDE │ STATUS │ PRICE │ AVG PRICE │ QUANTITY │ EXECUTED QTY │ CUM QUOTE │ TIF │        TIME         │     UPDATE TIME     │
├─────────────┼─────────┼──────┼────────┼───────────────┼────────┼───────┼───────────┼──────────┼──────────────┼───────────┼─────┼─────────────────────┼─────────────────────┤
│ 12803994759 │ ETHUSDT │ BUY  │ MARKET │ BOTH          │ FILLED │ 0     │ 2928.5    │ 0.003    │ 0.003        │ 8.7855    │ GTC │ 2025-12-16 09:37:12 │ 2025-12-16 09:37:12 │
│ 12929613862 │ ETHUSDT │ SELL │ LIMIT  │ BOTH          │ NEW    │ 4000  │ 0         │ 0.003    │ 0            │ 0         │ GTC │ 2025-12-18 05:26:45 │ 2025-12-18 05:26:45 │
└─────────────┴─────────┴──────┴────────┴───────────────┴────────┴───────┴───────────┴──────────┴──────────────┴───────────┴─────┴─────────────────────┴─────────────────────┘
```

### Order - Get Order
Query a single order by orderId or origClientOrderId.
```shell
./aster-cli futures order get --symbol=ETHUSDT --orderId=xxx
```
Or:
```shell
./aster-cli futures order get --symbol=ETHUSDT --origClientOrderId=xxx
```

### Order - Query trades
Exec: `./aster-cli futures order trade --symbol=BTCUSDT`

**Supported parameters:**
- `--symbol, -s`: Trading pair symbol (required)
- `--startTime, -a`: Start time (unix ms or "YYYY-MM-DD HH:MM:SS")
- `--endTime, -e`: End time (unix ms or "YYYY-MM-DD HH:MM:SS")
- `--fromId, -f`: Trade ID to fetch from
- `--limit, -l`: Number of results (default 500, max 1000)

```shell
┌─────────────┬─────────┬──────┬────────┬───────────────┬────────┬───────┬───────────┬──────────┬──────────────┬───────────┬─────┬─────────────────────┬─────────────────────┐
│  ORDER ID   │ SYMBOL  │ SIDE │  TYPE  │ POSITION SIDE │ STATUS │ PRICE │ AVG PRICE │ QUANTITY │ EXECUTED QTY │ CUM QUOTE │ TIF │        TIME         │     UPDATE TIME     │
├─────────────┼─────────┼──────┼────────┼───────────────┼────────┼───────┼───────────┼──────────┼──────────────┼───────────┼─────┼─────────────────────┼─────────────────────┤
│ 12803994759 │ ETHUSDT │ BUY  │ MARKET │ BOTH          │ FILLED │ 0     │ 2928.5    │ 0.003    │ 0.003        │ 8.7855    │ GTC │ 2025-12-16 09:37:12 │ 2025-12-16 09:37:12 │
└─────────────┴─────────┴──────┴────────┴───────────────┴────────┴───────┴───────────┴──────────┴──────────────┴───────────┴─────┴─────────────────────┴─────────────────────┘
```

### Order - Query force orders
Query user's force orders (liquidation orders).
```shell
./aster-cli futures order force
```

**Supported parameters:**
- `--symbol, -s`: Trading pair symbol
- `--autoCloseType, -t`: Auto close type: LIQUIDATION or ADL
- `--startTime, -a`: Start time (unix ms or "YYYY-MM-DD HH:MM:SS")
- `--endTime, -e`: End time (unix ms or "YYYY-MM-DD HH:MM:SS")
- `--limit, -l`: Number of results (default 50, max 100)

### Order - Cancel order by ID
Exec: `./aster-cli futures order cancel --symbol=ETHUSDT --orderId=xxx`

Or by client order ID:
```shell
./aster-cli futures order cancel --symbol=ETHUSDT --origClientOrderId=xxx
```

### Order - Cancel all orders by symbol
Exec: `./aster-cli futures order cancel --symbol=ETHUSDT`

If neither `--orderId` nor `--origClientOrderId` is provided, all open orders for the symbol will be canceled.

## Position

### Position - List positions
Exec: `./aster-cli futures position list`
```shell
┌─────────┬───────────────┬─────────────────┬─────────────┬───────────────────┬──────────┬─────────────────────┐
│ SYMBOL  │ POSITION SIDE │ POSITION AMOUNT │ ENTRY PRICE │ UNREALIZED PROFIT │ LEVERAGE │     UPDATE TIME     │
├─────────┼───────────────┼─────────────────┼─────────────┼───────────────────┼──────────┼─────────────────────┤
│ ETHUSDT │ BOTH          │ 0.003           │ 2928.5      │ -0.30615          │ 5        │ 2025-12-16 09:37:12 │
└─────────┴───────────────┴─────────────────┴─────────────┴───────────────────┴──────────┴─────────────────────┘
```

### Position - Show position risk
Exec: `./aster-cli futures position risk`

Or for a specific symbol:
```shell
./aster-cli futures position risk --symbol=BTCUSDT
```
```shell
┌─────────┬───────────────┬─────────────────┬──────────┬─────────────┬────────────┬───────────────────┬───────────────────┬─────────────────────┐
│ SYMBOL  │ POSITION SIDE │ POSITION AMOUNT │ NOTIONAL │ ENTRY PRICE │ MARK PRICE │ UNREALIZED PROFIT │ LIQUIDATION PRICE │     UPDATE TIME     │
├─────────┼───────────────┼─────────────────┼──────────┼─────────────┼────────────┼───────────────────┼───────────────────┼─────────────────────┤
│ ETHUSDT │ BOTH          │ 0.003           │ 0        │ 2928.5      │ 2823.95    │ -0.31365          │ 31.15841161       │ 2025-12-16 09:37:12 │
└─────────┴───────────────┴─────────────────┴──────────┴─────────────┴────────────┴───────────────────┴───────────────────┴─────────────────────┘
```

### Position - Position mode
**Get current position mode:**
```shell
./aster-cli futures position mode get
```

**Set position mode (Hedge/One-way):**
```shell
./aster-cli futures position mode set --dualSidePosition=true   # Hedge mode (dual side position)
./aster-cli futures position mode set --dualSidePosition=false  # One-way mode (single side position)
```

**Legacy commands (still supported):**
```shell
./aster-cli futures position side               # Get position mode status
./aster-cli futures position set-side --dualSidePosition=true  # Change position mode
```

### Position - Set position margin
Modify isolated position margin.
```shell
./aster-cli futures position set-margin --symbol=BTCUSDT --amount=1.0 --positionSide=SHORT --type=ADD
```

**Supported parameters:**
- `--symbol, -s`: Trading pair symbol (required)
- `--positionSide, -p`: Position side: BOTH, LONG, or SHORT (default BOTH)
- `--amount, -a`: Margin amount (required)
- `--type, -t`: Margin type: ADD or REDUCE (default ADD)

### Position - Margin history
Query position margin change history.
```shell
./aster-cli futures position margin-history --symbol=BTCUSDT
```

**Supported parameters:**
- `--symbol, -s`: Trading pair symbol (required)
- `--type, -t`: Margin type: 1 for Add, 2 for Reduce
- `--startTime, -a`: Start time (unix ms or "YYYY-MM-DD HH:MM:SS")
- `--endTime, -e`: End time (unix ms or "YYYY-MM-DD HH:MM:SS")
- `--limit, -l`: Number of results (default 500)

### Position - ADL quantile
Query ADL (Auto-Deleveraging) quantile estimation for positions.
```shell
./aster-cli futures position adl-quantile
```

Or for a specific symbol:
```shell
./aster-cli futures position adl-quantile --symbol=BTCUSDT
```

## Symbol

### Symbol - Set leverage
Change initial leverage for a symbol.
```shell
./aster-cli futures symbol set-leverage --symbol=BTCUSDT --leverage=10
```

### Symbol - Set margin type
Change symbol level margin type.
```shell
./aster-cli futures symbol set-margin-type --symbol=BTCUSDT --marginType=CROSSED
```
Supported types: ISOLATED, CROSSED

### Symbol - Leverage bracket
Query notional and leverage bracket information.
```shell
./aster-cli futures symbol leverage-bracket
```

Or for a specific symbol:
```shell
./aster-cli futures symbol leverage-bracket --symbol=BTCUSDT
```

## Funding

### Funding - Funding info
Query funding info including funding interval and fee cap/floor.
```shell
./aster-cli futures funding info
```

Or for a specific symbol:
```shell
./aster-cli futures funding info --symbol=BTCUSDT
```

### Funding - Funding rate history
Query funding rate history.
```shell
./aster-cli futures funding rate --symbol=BTCUSDT
```

**Supported parameters:**
- `--symbol, -s`: Trading pair symbol
- `--startTime, -a`: Start time (unix ms or "YYYY-MM-DD HH:MM:SS")
- `--endTime, -e`: End time (unix ms or "YYYY-MM-DD HH:MM:SS")
- `--limit, -l`: Number of results (default 100, max 1000)
