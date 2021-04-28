package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spacemonkeygo/openssl"
	cli "github.com/urfave/cli/v2"

	sizzle "github.com/wirasuta/sizzle/web-owner-client/abi"
	"github.com/wirasuta/sizzle/web-owner-client/utils"
)

var config *utils.Config

func createCertificate(country string, organization string, domain string) error {
	rsaKey, err := openssl.GenerateRSAKey(1024)
	if err != nil {
		log.Fatal(err)
	}
	certInfo := &openssl.CertificateInfo{
		Serial:       big.NewInt(int64(1)),
		Country:      country,
		Organization: organization,
		CommonName:   domain,
		Issued:       0,
		Expires:      time.Hour * 24 * 90,
	}
	cert, err := openssl.NewCertificate(certInfo, rsaKey)
	if err != nil {
		log.Fatal(err)
	}
	if err != cert.Sign(rsaKey, openssl.EVP_SHA256) {
		log.Fatal(err)
	}
	certPem, err := cert.MarshalPEM()
	if err != nil {
		log.Fatal(err)
	}
	keyPem, err := rsaKey.MarshalPKCS1PrivateKeyPEM()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(keyPem))
	fmt.Print(string(certPem))

	return nil
}

func publishCertificate(cert *openssl.Certificate, certPrivateKey *openssl.PrivateKey) error {
	pubKey, err := cert.PublicKey()
	if err != nil {
		log.Fatal(err)
	}
	pubKeyPem, err := pubKey.MarshalPKIXPublicKeyPEM()
	if err != nil {
		log.Fatal(err)
	}
	subject, err := cert.GetSubjectName()
	if err != nil {
		log.Fatal(err)
	}
	domain, ok := subject.GetEntry(openssl.NID_commonName)
	if !ok {
		log.Fatal("Fail to get domain from certificate common name")
	}

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
	sizzleAddress := common.HexToAddress(config.SizzleAddress)
	client, err := ethclient.Dial(config.EthClientUrl)
	if err != nil {
		log.Fatal(err)
	}
	szlTxr, err := sizzle.NewSizzleTransactor(sizzleAddress, client)
	if err != nil {
		log.Fatal(err)
	}
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
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	_, err = szlTxr.CertPublishRequest(auth, domain, string(pubKeyPem))
	if err != nil {
		log.Fatal(err)
	}

	signedSzlAddress, err := (*certPrivateKey).SignPKCS1v15(openssl.SHA256_Method, []byte(config.SizzleAddress))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Please add the following entry to %s DNS TXT record: \"sig=%s\"", domain, hex.EncodeToString(signedSzlAddress))

	return nil
}

func handleCreate(ctx *cli.Context) error {
	return createCertificate(ctx.String("country"), ctx.String("organization"), ctx.String("domain"))
}

func publishSubmit(ctx *cli.Context) error {
	certByte, err := ioutil.ReadFile(ctx.Path("cert"))
	if err != nil {
		log.Fatal(err)
	}
	privateKeyByte, err := ioutil.ReadFile(ctx.Path("privatekey"))
	if err != nil {
		log.Fatal(err)
	}
	cert, err := openssl.LoadCertificateFromPEM(certByte)
	if err != nil {
		log.Println("here")
		log.Fatal(err)
	}
	privateKey, err := openssl.LoadPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		log.Fatal(err)
	}

	publishCertificate(cert, &privateKey)

	return nil
}

func handleInit(ctx *cli.Context) error {
	// TODO: Chain generate and submit
	return nil
}

func main() {
	config = utils.LoadConfig()
	app := &cli.App{
		Name:  "sizzle-web-owner",
		Usage: "Sizzle CLI for web owner",
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Create self-signed TLS certificate",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "country",
						Aliases:  []string{"u"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "organization",
						Aliases:  []string{"o"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "domain",
						Aliases:  []string{"d"},
						Required: true,
					},
				},
				Action: handleCreate,
			},
			{
				Name:  "publish",
				Usage: "Publish created self-signed TLS certificate",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:    "privatekey",
						Aliases: []string{"k"},
					},
					&cli.PathFlag{
						Name:    "cert",
						Aliases: []string{"c"},
					},
				},
				Action: publishSubmit,
			},
			{
				Name:  "init",
				Usage: "Generate and submit self-signed TLS certificate",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "domain",
						Aliases: []string{"d"},
					},
				},
				Action: handleInit,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
