package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spacemonkeygo/openssl"
	"github.com/urfave/cli/v2"

	sizzle "github.com/wirasuta/sizzle/end-user-client/abi"
	"github.com/wirasuta/sizzle/end-user-client/utils"
)

var config *utils.Config

func loadAndVerify(cert *openssl.Certificate, certPath string) error {
	subject, err := cert.GetSubjectName()
	if err != nil {
		log.Fatal(err)
	}
	domain, ok := subject.GetEntry(openssl.NID_commonName)
	if !ok {
		log.Fatal("Fail to get domain from certificate common name")
	}
	certPresent := false
	certUtil := utils.CertUtil{
		NssDatabaseDirectory: config.NssDatabaseDirectory,
	}
	res, err := certUtil.List()
	if err != nil {
		log.Fatal(err)
	}
	for _, cert := range res {
		if cert.Nickname == domain {
			certPresent = true
		}
	}

	if !certPresent {
		sizzleAddress := common.HexToAddress(config.SizzleAddress)
		client, err := ethclient.Dial(config.EthClientUrl)
		if err != nil {
			log.Fatal(err)
		}
		szlClr, err := sizzle.NewSizzleCaller(sizzleAddress, client)
		if err != nil {
			log.Fatal(err)
		}
		certMetadata, err := szlClr.CertQuery(&bind.CallOpts{}, domain)
		if err != nil {
			log.Fatal(err)
		}

		if certMetadata.Valid {
			err = certUtil.AddSSL(domain, certPath)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Validated and added %s certificate\n", domain)
		}
	}

	return nil
}

func endorse(domain string) error {
	privateKey := utils.LoadPrivateKey(config)
	sizzleAddress := common.HexToAddress(config.SizzleAddress)
	client, err := ethclient.Dial(config.EthClientUrl)
	if err != nil {
		log.Fatal(err)
	}
	szlTxr, err := sizzle.NewSizzleTransactor(sizzleAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	auth := utils.GenerateAuthBind(privateKey, client, 125000)

	_, err = szlTxr.CertEndorseByUser(auth, domain)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func deny(domain string) error {
	privateKey := utils.LoadPrivateKey(config)
	sizzleAddress := common.HexToAddress(config.SizzleAddress)
	client, err := ethclient.Dial(config.EthClientUrl)
	if err != nil {
		log.Fatal(err)
	}
	szlTxr, err := sizzle.NewSizzleTransactor(sizzleAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	auth := utils.GenerateAuthBind(privateKey, client, 125000)

	_, err = szlTxr.CertDenyByUser(auth, domain)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func handleVerify(ctx *cli.Context) error {
	certPath := ctx.Path("cert")
	certByte, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Fatal(err)
	}
	cert, err := openssl.LoadCertificateFromPEM(certByte)
	if err != nil {
		log.Fatal(err)
	}

	return loadAndVerify(cert, certPath)
}

func handleEndorse(ctx *cli.Context) error {
	domain := ctx.String("domain")
	return endorse(domain)
}

func handleDeny(ctx *cli.Context) error {
	domain := ctx.String("domain")
	return deny(domain)
}

func main() {
	config = utils.LoadConfig()
	app := &cli.App{
		Name:  "sizzle-end-user",
		Usage: "Sizzle CLI for end user",
		Commands: []*cli.Command{
			{
				Name:  "verify",
				Usage: "Verify and load TLS certificate",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "cert",
						Aliases:  []string{"c"},
						Required: true,
					},
				},
				Action: handleVerify,
			},
			{
				Name:  "endorse",
				Usage: "Endorse a published certificate",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "domain",
						Aliases:  []string{"d"},
						Required: true,
					},
				},
				Action: handleEndorse,
			},
			{
				Name:  "deny",
				Usage: "Deny a published certificate",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "domain",
						Aliases:  []string{"d"},
						Required: true,
					},
				},
				Action: handleDeny,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
