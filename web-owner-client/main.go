package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spacemonkeygo/openssl"
	cli "github.com/urfave/cli/v2"

	sizzle "github.com/wirasuta/sizzle/web-owner-client/abi"
	"github.com/wirasuta/sizzle/web-owner-client/utils"
)

var config *utils.Config
var verbose bool

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

	_, err = szlTxr.CertPublishRequest(auth, domain, string(pubKeyPem))
	if err != nil {
		log.Fatal(err)
	}

	signedSzlAddress, err := (*certPrivateKey).SignPKCS1v15(openssl.SHA256_Method, []byte(config.SizzleAddress))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Please add the following entry to %s DNS TXT record: \"sig=%s\"\n", domain, hex.EncodeToString(signedSzlAddress))

	return nil
}

func rekeyCertificate(cert *openssl.Certificate, certPrivateKey *openssl.PrivateKey) error {
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
	auth := utils.GenerateAuthBind(privateKey, client, 100000)

	_, err = szlTxr.CertRekey(auth, domain, string(pubKeyPem))
	if err != nil {
		log.Fatal(err)
	}

	signedSzlAddress, err := (*certPrivateKey).SignPKCS1v15(openssl.SHA256_Method, []byte(config.SizzleAddress))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Please add the following entry to %s DNS TXT record: \"sig=%s\"\n", domain, hex.EncodeToString(signedSzlAddress))

	return nil
}

func revokeCertificate(domain string) error {
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
	auth := utils.GenerateAuthBind(privateKey, client, 60000)

	_, err = szlTxr.CertRevoke(auth, domain)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func handleCreate(ctx *cli.Context) error {
	verbose = ctx.Bool("verbose")
	utils.VerboseTime("handleCreate", verbose)
	err := createCertificate(ctx.String("country"), ctx.String("organization"), ctx.String("domain"))
	utils.VerboseTime("handlePublish end", verbose)
	return err
}

func handlePublish(ctx *cli.Context) error {
	verbose = ctx.Bool("verbose")
	utils.VerboseTime("handlePublish", verbose)
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
	utils.VerboseTime("handlePublish end", verbose)
	return nil
}

func handleRekey(ctx *cli.Context) error {
	verbose = ctx.Bool("verbose")
	utils.VerboseTime("handleRekey", verbose)
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

	rekeyCertificate(cert, &privateKey)
	utils.VerboseTime("handleRekey end", verbose)
	return nil
}

func handleRevoke(ctx *cli.Context) error {
	verbose = ctx.Bool("verbose")
	utils.VerboseTime("handleRevoke", verbose)
	err := revokeCertificate(ctx.String("domain"))
	utils.VerboseTime("handleRevoke end", verbose)
	return err
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
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
			},
		},
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
				Action: handlePublish,
			},
			{
				Name:  "rekey",
				Usage: "Rekey published TLS certificate",
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
				Action: handleRekey,
			},
			{
				Name:  "revoke",
				Usage: "Revoke published TLS certificate",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:    "domain",
						Aliases: []string{"d"},
					},
				},
				Action: handleRevoke,
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
