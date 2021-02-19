package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func GenerateTLSConfig(certFile, keyFile string) (*tls.Config, error) {
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	if len(certificate.Certificate) < 1 || certificate.PrivateKey == nil {
		return nil, fmt.Errorf("unable to import certificate, length is too small")
	}

	return &tls.Config{Certificates: []tls.Certificate{certificate}}, nil
}

type TLSCertificateFields struct {
	Country            string
	StateName          string
	Locality           string
	Organization       string
	OrganizationalUnit string
	CommonName         string
	Expiration         time.Time
}

func GenerateSelfSignedTLSCertificate(fields *TLSCertificateFields, certFile, keyFile string) error {
	template := &x509.Certificate{
		IsCA:                  true,
		BasicConstraintsValid: true,
		SubjectKeyId:          []byte{1, 2, 3},
		SerialNumber:          big.NewInt(1234),
		NotBefore:             time.Now(),
		NotAfter:              fields.Expiration,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	// Fill optional fields

	if fields.Country != "" {
		template.Subject.Country = []string{fields.Country}
	}
	if fields.StateName != "" {
		template.Subject.Province = []string{fields.StateName}
	}
	if fields.Locality != "" {
		template.Subject.Locality = []string{fields.Locality}
	}
	if fields.Organization != "" {
		template.Subject.Organization = []string{fields.Organization}
	}
	if fields.OrganizationalUnit != "" {
		template.Subject.OrganizationalUnit = []string{fields.OrganizationalUnit}
	}
	if fields.CommonName != "" {
		template.Subject.CommonName = fields.CommonName
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	publicKey := &privateKey.PublicKey

	cert, err := x509.CreateCertificate(rand.Reader, template, template, publicKey, privateKey)
	if err != nil {
		return err
	}

	// Write certificate
	certFileP, err := os.Create(certFile)
	if err != nil {
		return err
	}
	defer certFileP.Close()

	pemBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	}
	err = pem.Encode(certFileP, pemBlock)
	if err != nil {
		return err
	}

	// Write private key
	keyFileP, err := os.Create(keyFile)
	if err != nil {
		return err
	}
	defer keyFileP.Close()

	pemBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	return pem.Encode(keyFileP, pemBlock)
}
