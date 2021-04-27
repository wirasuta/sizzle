package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EthClientUrl      string
	EthPrivateKey     string
	EthPrivateKeyFile string
	SizzleAddress     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	conf := &Config{
		EthClientUrl:      os.Getenv("ETH_CLIENT_URL"),
		EthPrivateKeyFile: os.Getenv("ETH_PRIVATE_KEY_FILE"),
		EthPrivateKey:     os.Getenv("ETH_PRIVATE_KEY"),
		SizzleAddress:     os.Getenv("SIZZLE_ADDRESS"),
	}

	return conf
}
