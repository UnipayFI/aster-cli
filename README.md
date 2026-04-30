# Aster CLI

A command-line tool for AsterDEX API developed in Go, supporting spot and futures trading functions.

Built on the Aster V3 API (EIP-712 + ECDSA "API wallet" signing). The legacy V1
HMAC scheme is no longer supported.

## Installation and Configuration

### Installation
```shell
curl -sSL https://raw.githubusercontent.com/UnipayFI/aster-cli/refs/heads/main/download.sh | bash
```

### Environment variables
Before using, set your AsterDEX V3 API-wallet credentials:
```shell
export API_ADDRESS="0x..."        # main wallet address
export API_PRIVATE_KEY="0x..."    # API wallet private key (hex)
# Optional. Defaults to mainnet (1666). Set to 714 for testnet.
export API_CHAIN_ID="1666"
```
`API_ADDRESS` is the master wallet address; its private key never touches the
client. `API_PRIVATE_KEY` is the API wallet's private key — the SDK derives the
signer address from it.

### Output format
Every command supports a global `--json` flag. Without it, results render as a
table; with it, the raw API response is printed as indented JSON, e.g.:
```shell
./aster-cli spot account --json
```

## Usage
All commands are to be used in the following format:
```
./aster-cli [Module] [Subcommand] [Arguments]

Available Commands:
  futures     Futures trading commands
  help        Help about any command
  spot        Spot trading commands
  wallet      Wallet commands
```

Each leaf subcommand's `-h` output includes a `Docs Link:` pointing to the
official Aster API documentation page for that endpoint.

### Spot Module
Exec: `./aster-cli spot [Subcommand] [Arguments]`
```shell
Available Commands:
  account          Show account info
  balance          Show account balances (non-zero only)
  commission-rate  Get commission rate for a symbol
  order            Support create, cancel, list, get orders
  trade            Query user trades
```
**[View detailed documentation](docs/spot.md)**


### Futures Module
Exec: `./aster-cli futures [Subcommand] [Arguments]`
```shell
Available Commands:
  account     Account management (balances, info, commission-rate, income, multi-assets-mode)
  funding     Funding rate commands (info, rate history)
  order       Order management (create, cancel, list, get, trade, force)
  position    Position management (list, risk, mode, margin, margin-history, adl-quantile)
  symbol      Symbol configuration (leverage, margin-type, leverage-bracket)
```
**[View detailed documentation](docs/futures.md)**

### Wallet Module
Exec: `./aster-cli wallet [Subcommand] [Arguments]`
```shell
Available Commands:
  transfer    Transfer assets between spot and futures
```
**[View detailed documentation](docs/wallet.md)**

