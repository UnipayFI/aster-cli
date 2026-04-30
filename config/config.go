package config

import (
	"os"
	"strconv"
)

var Config struct {
	APIAddress       string
	APIPrivateKey    string
	ChainID          int64
	SpotBaseURL      string
	FuturesBaseURL   string
	OutputJSON       bool
}

func init() {
	Config.APIAddress = os.Getenv("API_ADDRESS")
	Config.APIPrivateKey = os.Getenv("API_PRIVATE_KEY")
	Config.SpotBaseURL = os.Getenv("API_SPOT_BASE_URL")
	Config.FuturesBaseURL = os.Getenv("API_FUTURES_BASE_URL")
	if v := os.Getenv("API_CHAIN_ID"); v != "" {
		if id, err := strconv.ParseInt(v, 10, 64); err == nil {
			Config.ChainID = id
		}
	}
}
