package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type CertUtil struct {
	NssDatabaseDirectory string
}

const (
	Prohibited      = "p"
	TrustedPeer     = "P"
	ValidCA         = "c"
	TrustedCAClient = "T"
	TrustedCAServer = "C"
	UserCert        = "u"
	SendWarning     = "w"
	None            = ""
)

type CertEntry struct {
	Nickname string
	SSL      string
	SMIME    string
	JARXPI   string
}

func parseCertUtilLine(cert string) *CertEntry {
	nickname, rest := cert[:61], cert[61:]
	nickname = strings.TrimRight(nickname, " ")
	trustAttrs := strings.Split(rest, ",")
	jarxpiTrustAttrs := trustAttrs[len(trustAttrs)-1]
	smimeTrustAttrs := trustAttrs[len(trustAttrs)-2]
	sslTrustAttrs := trustAttrs[len(trustAttrs)-3]
	if len(sslTrustAttrs) > 0 {
		sslTrustAttrs = string(sslTrustAttrs[len(sslTrustAttrs)-1])
	}
	if sslTrustAttrs == " " {
		sslTrustAttrs = ""
	}
	certEntry := &CertEntry{
		Nickname: nickname,
		SSL:      sslTrustAttrs,
		SMIME:    smimeTrustAttrs,
		JARXPI:   jarxpiTrustAttrs,
	}

	return certEntry
}

func (certutil CertUtil) List() ([]*CertEntry, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	out, err := exec.Command("certutil", "-L", "-d", path.Join(home, certutil.NssDatabaseDirectory)).Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(out), "\n")
	lines = lines[4:]
	certEntries := []*CertEntry{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		certEntries = append(certEntries, parseCertUtilLine(line))
	}

	return certEntries, nil
}

func (certutil CertUtil) AddSSL(nickname string, certPath string, trustAttrs string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command("certutil", "-A", "-n", nickname, "-i", certPath, "-t", trustAttrs, "-d", path.Join(home, certutil.NssDatabaseDirectory)).Run()
	if err != nil {
		log.Println(nickname)
		log.Println(certPath)
		log.Fatal(err)
	}

	return nil
}

func GenerateTrustAttrs(SSL string, SMIME string, JARXPI string) string {
	return fmt.Sprintf("%s,%s,%s", SSL, SMIME, JARXPI)
}
