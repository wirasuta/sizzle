package main

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spacemonkeygo/openssl"
	cli "github.com/urfave/cli/v2"

	sizzle "github.com/wirasuta/sizzle/peer-client/abi"
	"github.com/wirasuta/sizzle/peer-client/utils"
)

var config *utils.Config

func register() error {
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
	auth := utils.GenerateAuthBind(privateKey, client, 50000)

	_, err = szlTxr.PeerRegister(auth)
	if err != nil {
		log.Fatal(err)
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

	_, err = szlTxr.CertEndorseByPeer(auth, domain)
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

	_, err = szlTxr.CertDenyByPeer(auth, domain)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func verify(domain string, pubKey openssl.PublicKey) (bool, error) {
	txtrec, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatal(err)
	}

	verified := false
	for _, txt := range txtrec {
		if strings.HasPrefix(txt, "sig=") {
			sig := strings.TrimPrefix(txt, "sig=")
			sigBytes, err := hex.DecodeString(sig)
			if err != nil {
				log.Fatal(err)
			}
			err = pubKey.VerifyPKCS1v15(openssl.SHA256_Method, []byte(config.SizzleAddress), sigBytes)
			if err == nil {
				log.Printf("Verified %s!\n", domain)
				verified = true
			} else {
				log.Println(sig)
				log.Println(config.SizzleAddress)
				log.Println(err)
			}
		}
	}

	return verified, nil
}

func listen() error {
	sizzleAddress := common.HexToAddress(config.SizzleAddress)
	client, err := ethclient.Dial(config.EthClientUrl)
	if err != nil {
		log.Fatal(err)
	}
	szlFilterer, err := sizzle.NewSizzleFilterer(sizzleAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	certPublishedChan := make(chan *sizzle.SizzleCertPublishRequestCreated)
	_, err = szlFilterer.WatchCertPublishRequestCreated(&bind.WatchOpts{}, certPublishedChan)
	if err != nil {
		log.Fatal(err)
	}

	for cert := range certPublishedChan {
		domain := cert.Domain
		publicKey, err := openssl.LoadPublicKeyFromPEM([]byte(cert.PubKey))
		if err != nil {
			log.Fatal(err)
		}

		ok, err := verify(domain, publicKey)
		if err != nil {
			log.Fatal(err)
		}

		if !ok {
			// TODO: Implement retry with exponential backoff
			log.Printf("Failed to verify %s ownership\n", domain)
			deny(domain)
		} else {
			endorse(domain)
		}
	}

	return nil
}

func handleRegister(ctx *cli.Context) error {
	return register()
}

func handleEndorse(ctx *cli.Context) error {
	domain := ctx.String("domain")
	return endorse(domain)
}

func handleDeny(ctx *cli.Context) error {
	domain := ctx.String("domain")
	return deny(domain)
}

func handleVerify(ctx *cli.Context) error {
	domain := ctx.String("domain")
	pubKeyByte, err := ioutil.ReadFile(ctx.Path("publickey"))
	if err != nil {
		log.Fatal(err)
	}
	pubKey, err := openssl.LoadPublicKeyFromPEM(pubKeyByte)
	if err != nil {
		log.Fatal(pubKey)
	}

	success, err := verify(domain, pubKey)
	if err != nil {
		log.Fatal(err)
	}

	if success != true {
		log.Fatal("Validation failed")
	}

	return nil
}

func handleListen(ctx *cli.Context) error {
	return listen()
}

func main() {
	config = utils.LoadConfig()
	app := &cli.App{
		Name:  "sizzle-peer",
		Usage: "Sizzle CLI for verifying peer",
		Commands: []*cli.Command{
			{
				Name:   "register",
				Usage:  "Register your address as a verifier for sizzle",
				Action: handleRegister,
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
			{
				Name:  "verify",
				Usage: "Verify certificate using DNS record",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "domain",
						Aliases:  []string{"d"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "publickey",
						Aliases:  []string{"p"},
						Required: true,
					},
				},
				Action: handleVerify,
			},
			{
				Name:   "listen",
				Usage:  "Listen for published certificate, verify, and endorse/deny accordingly",
				Action: handleListen,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
