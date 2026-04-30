# Wallet Module

> The global `--json` flag is supported on every command and prints the raw
> API response as indented JSON instead of a table.

## Quick Navigation
- [Transfer](#transfer---transfer-assets-between-spot-and-futures)

## Transfer - Transfer assets between spot and futures
Docs Link: <https://asterdex.github.io/aster-api-website/futures-v3/account%26trades/#transfer-between-futures-and-spot-transfer>

Transfer assets between spot and futures wallets.

```shell
./aster-cli wallet transfer --kindType=SPOT_FUTURE --asset=USDT --amount=100
```

**Supported transfer types:**
- `SPOT_FUTURE`: Transfer from spot wallet to futures wallet
- `FUTURE_SPOT`: Transfer from futures wallet to spot wallet

**Parameters:**
- `--kindType, -t`: Transfer type: SPOT_FUTURE or FUTURE_SPOT (required)
- `--asset, -a`: Asset to transfer, e.g., USDT, BTC (required)
- `--amount, -m`: Amount to transfer (decimal string, must be greater than 0)

**Examples:**

Transfer 100 USDT from spot to futures:
```shell
./aster-cli wallet transfer --kindType=SPOT_FUTURE --asset=USDT --amount=100
```

Transfer 50 USDC from futures to spot:
```shell
./aster-cli wallet transfer --kindType=FUTURE_SPOT --asset=USDC --amount=50
```

**Output:**
```shell
┌────────────────┬─────────┐
│ TRANSACTION ID │ STATUS  │
├────────────────┼─────────┤
│ 50548639       │ SUCCESS │
└────────────────┴─────────┘
```

With `--json`:
```json
{
  "tranId": 50548639,
  "status": "SUCCESS"
}
```
