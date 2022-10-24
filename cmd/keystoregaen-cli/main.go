package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pojntfx/keystoregaen/pkg/utils"
)

func main() {
	storepass := flag.String("storepass", "", "Password of the keystore to generate")
	keypass := flag.String("keypass", "", "Password of the certificate to generate")
	alias := flag.String("alias", "", "Alias of the certificate to generate")
	cname := flag.String("cname", "", "CNAME of the certificate to generate")
	validity := flag.Duration("validity", time.Hour*24*365, "Validity of certificate to generate")
	bits := flag.Uint("bits", 2048, "Bits to use to generate RSA key")
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

	out, err := os.Create(*dst)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	log.Println("Generating keystore ...")

	if err := utils.GenerateKeystore(
		*storepass,
		*keypass,
		*alias,
		*cname,
		*validity,
		uint32(*bits),
		out,
	); err != nil {
		panic(err)
	}
}
