# Wallet Module

## Quick Navigation
- [Transfer](#transfer---transfer-assets-between-spot-and-futures)

## Transfer - Transfer assets between spot and futures

Transfer assets between spot and futures wallets.

```shell
./aster-cli wallet transfer --type=SPOT_FUTURE --asset=USDT --amount=100
```

**Supported transfer types:**
- `SPOT_FUTURE`: Transfer from spot wallet to futures wallet
- `FUTURE_SPOT`: Transfer from futures wallet to spot wallet

**Parameters:**
- `--type, -t`: Transfer type: SPOT_FUTURE or FUTURE_SPOT (required)
- `--asset, -a`: Asset to transfer, e.g., USDT, BTC (required)
- `--amount, -m`: Amount to transfer (required, must be greater than 0)

**Examples:**

Transfer 100 USDT from spot to futures:
```shell
./aster-cli wallet transfer --type=SPOT_FUTURE --asset=USDT --amount=100
```

Transfer 50 USDC from futures to spot:
```shell
./aster-cli wallet transfer --type=FUTURE_SPOT --asset=USDC --amount=50
```

**Output:**
```shell
Transfer successful!
┌────────────────┬─────────┐
│ TRANSACTION ID │ STATUS  │
├────────────────┼─────────┤
│ 50548639       │ SUCCESS │
└────────────────┴─────────┘
```
