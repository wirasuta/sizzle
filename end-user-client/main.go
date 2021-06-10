package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
var verbose bool

func loadAndVerify(domain string) error {
	tmpCert, err := ioutil.TempFile(os.TempDir(), "cert-*.pem")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpCert.Name())
	certDomain := fmt.Sprintf("http://cert.%s", domain)
	resp, err := http.Get(certDomain)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	certPem, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	cert, err := openssl.LoadCertificateFromPEM(certPem)
	if err != nil {
		log.Fatal(err)
	}
	_, err = tmpCert.Write(certPem)
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

		if certMetadata.Status == sizzle.CertStatusValid {
			trustAttrs := utils.GenerateTrustAttrs(utils.TrustedCAServer, utils.None, utils.None)
			err = certUtil.AddSSL(domain, tmpCert.Name(), trustAttrs)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Validated and added %s certificate\n", domain)
		} else if certMetadata.Status == sizzle.CertStatusInvalid {
			log.Printf("%s certificate is invalid\n", domain)
		} else if certMetadata.Status == sizzle.CertStatusRevoked {
			log.Printf("%s certificate has been revoked\n", domain)
		} else {
			log.Printf("%s certificate unknown\n", domain)
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
	auth := utils.GenerateAuthBind(privateKey, client, 500000)

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
	auth := utils.GenerateAuthBind(privateKey, client, 500000)

	_, err = szlTxr.CertDenyByUser(auth, domain)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func handleVerify(ctx *cli.Context) error {
	verbose = ctx.Bool("verbose")
	utils.VerboseTime("handleVerify", verbose)
	domain := ctx.Path("domain")
	err := loadAndVerify(domain)
	utils.VerboseTime("handleVerify end", verbose)
	return err
}

func handleEndorse(ctx *cli.Context) error {
	verbose = ctx.Bool("verbose")
	utils.VerboseTime("handleEndorse", verbose)
	domain := ctx.String("domain")
	err := endorse(domain)
	utils.VerboseTime("handleEndorse end", verbose)
	return err
}

func handleDeny(ctx *cli.Context) error {
	verbose = ctx.Bool("verbose")
	utils.VerboseTime("handleDeny", verbose)
	domain := ctx.String("domain")
	err := deny(domain)
	utils.VerboseTime("handleDeny end", verbose)
	return err
}

func main() {
	config = utils.LoadConfig()
	app := &cli.App{
		Name:  "sizzle-end-user",
		Usage: "Sizzle CLI for end user",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "verify",
				Usage: "Verify and load TLS certificate of a domain",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "domain",
						Aliases:  []string{"d"},
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
