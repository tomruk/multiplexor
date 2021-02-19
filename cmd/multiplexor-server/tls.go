package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/tomruk/multiplexor/utils"
)

var shell, _ = readline.New("")

func generateSelfSignedTLSCert() error {
	fields := new(utils.TLSCertificateFields)

	fmt.Println("Country:")
	fields.Country = readlineWithDefault("TR")

	fmt.Println("State or province:")
	fields.StateName = readlineWithDefault("somewhere")

	fmt.Println("Locality:")
	fields.Locality = readlineWithDefault("")

	fmt.Println("Organization:")
	fields.Organization = readlineWithDefault("")

	fmt.Println("Organizational unit:")
	fields.OrganizationalUnit = readlineWithDefault("")

	fmt.Println("Common name:")
	fields.CommonName = readlineWithDefault("")

	fmt.Println("Expiration (days)")
	expiration, err := strconv.Atoi(readlineWithDefault("365"))
	if err != nil {
		return fmt.Errorf("integer required")
	}

	fields.Expiration = time.Now().Add(time.Hour * 24 * time.Duration(expiration))
	fmt.Printf("Certificate expires at: %s\n\n", fields.Expiration.String())

	fmt.Println("Save certificate to")
	certFile := readlineWithDefault("cert.pem")

	fmt.Println("Save private key to")
	keyFile := readlineWithDefault("key.pem")

	// If certificate file exists
	if _, err := os.Stat(certFile); !os.IsNotExist(err) {
		overwrite(certFile)
	}

	// If private key file exists
	if _, err := os.Stat(keyFile); !os.IsNotExist(err) {
		overwrite(keyFile)
	}

	fmt.Println("Generating self signed TLS certificate")

	err = utils.GenerateSelfSignedTLSCertificate(fields, certFile, keyFile)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully saved certificate/private key to: %s and %s\n", certFile, keyFile)
	return nil
}

func overwrite(file string) {
	fmt.Printf("%s exists, overwrite? Y/n\n", file)
	yesNo := readlineWithDefault("y")
	yesNo = strings.TrimSpace(strings.ToLower(yesNo))

	if yesNo == "y" || yesNo == "" {
		return
	}
	fmt.Println("Cancelled")
	os.Exit(0)
}

func readlineWithDefault(defaultInput string) string {
	line, err := shell.ReadlineWithDefault(defaultInput)
	if err == readline.ErrInterrupt {
		fmt.Println("Cancelled")
		os.Exit(0)
	}
	return line
}
