package utils

import (
	"os"
)

type Config struct {
	EthClientUrl  string
	SizzleAddress string
}

func LoadConfig() *Config {
	conf := &Config{
		EthClientUrl:  os.Getenv("ETH_CLIENT_URL"),
		SizzleAddress: os.Getenv("SIZZLE_ADDRESS"),
	}

	return conf
}
