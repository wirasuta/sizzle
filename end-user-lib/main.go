package main

import "C"

import (
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	sizzle "github.com/wirasuta/sizzle/end-user-lib/abi"
	"github.com/wirasuta/sizzle/end-user-lib/utils"
)

//export VerifyBlockchainCert
func VerifyBlockchainCert(cdomain *C.char) bool {
	conf := utils.LoadConfig()
	sizzleAddress := common.HexToAddress(conf.SizzleAddress)
	client, err := ethclient.Dial(conf.EthClientUrl)
	if err != nil {
		log.Fatal(err)
	}
	szlClr, err := sizzle.NewSizzleCaller(sizzleAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	domain := strings.TrimRight(C.GoString(cdomain), "\x00\n ")
	certMetadata, err := szlClr.CertQuery(&bind.CallOpts{}, domain)
	if err != nil {
		log.Fatal(err)
	}

	if certMetadata.Status == sizzle.CertStatusValid {
		log.Printf("Validated %s certificate\n", domain)
		return true
	} else {
		return false
	}
}

func main() {
}
