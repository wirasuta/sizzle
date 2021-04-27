package utils

import (
	"context"
	"crypto/ecdsa"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func LoadPrivateKey(config *Config) *ecdsa.PrivateKey {
	var privateKeyStr string
	if len(config.EthPrivateKey) > 0 {
		privateKeyStr = config.EthPrivateKey
	} else {
		temp, err := ioutil.ReadFile(config.EthPrivateKeyFile)
		if err != nil {
			log.Fatal(err)
		}

		privateKeyStr = string(temp)
	}
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	return privateKey
}

func GenerateAuthBind(privateKey *ecdsa.PrivateKey, client *ethclient.Client, gasLimit uint64) *bind.TransactOpts {
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return auth
}
