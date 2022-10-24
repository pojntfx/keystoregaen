package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"flag"
	"math"
	"math/big"
	"os"
	"strings"
	"time"

	ks "github.com/pavlo-v-chernykh/keystore-go/v4"
)

func main() {
	storepass := flag.String("storepass", "", "Password of the keystore to generate")
	keypass := flag.String("keypass", "", "Password of the certificate to generate")
	alias := flag.String("alias", "", "Alias of the certificate to generate")
	cname := flag.String("cname", "", "CNAME of the certificate to generate")
	validity := flag.Duration("validity", time.Hour*24*365, "Validity of certificate to generate")
	bits := flag.Int("bits", 2048, "Bits to use to generate RSA key")
	dst := flag.String("dst", "keystoregaen.jks", "Path to write the keystore to")

	flag.Parse()

	if strings.TrimSpace(*storepass) == "" {
		panic("could not continue with empty storepass")
	}

	if strings.TrimSpace(*keypass) == "" {
		panic("could not continue with empty keypass")
	}

	if strings.TrimSpace(*alias) == "" {
		panic("could not continue with empty alias")
	}

	if strings.TrimSpace(*cname) == "" {
		panic("could not continue with empty CNAME")
	}

	// Generate private key
	key, err := rsa.GenerateKey(rand.Reader, *bits)
	if err != nil {
		panic(err)
	}

	rawKey, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		panic(err)
	}

	// Generate certificate
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}

	now := time.Now()
	tpl := &x509.Certificate{
		SerialNumber: serialNumber,
		NotBefore:    now,
		NotAfter:     now.Add(*validity),
		Subject: pkix.Name{
			CommonName: *cname,
		},
		Issuer: pkix.Name{
			CommonName: *cname,
		},
	}

	cert, err := x509.CreateCertificate(rand.Reader, tpl, tpl, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}

	// Generate & write keystore
	keystore := ks.New()

	if err := keystore.SetPrivateKeyEntry(
		*alias,
		ks.PrivateKeyEntry{
			CreationTime: time.Now(),
			PrivateKey:   rawKey,
			CertificateChain: []ks.Certificate{
				{
					Type:    "X509",
					Content: cert,
				},
			},
		},
		[]byte(*keypass),
	); err != nil {
		panic(err)
	}

	out, err := os.Create(*dst)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	if err := keystore.Store(out, []byte(*storepass)); err != nil {
		panic(err)
	}
}
