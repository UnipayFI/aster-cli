# Aster CLI

A command-line tool for AsterDEX API developed in Go, supporting spot and futures trading functions.

## Installation and Configuration

### Installation
```shell
curl -sSL https://raw.githubusercontent.com/UnipayFI/aster-cli/refs/heads/main/download.sh | bash
```

### Environment variables
Before using, you need to set the AsterDEX API key:
```shell
export API_KEY="your_api_key"
export API_SECRET="your_api_secret"
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

