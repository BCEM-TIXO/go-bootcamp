package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	CertFile = "../cert/client/cert.pem"
	KeyFile  = "../cert/client/key.pem"
	MiniCA   = "../cert/minica.pem"
)

type ClientFlags struct {
	k string
	c int64
	m int64
}

func init() {
	flag.StringVar(&fl.k, "k", "", "candy type")
	flag.Int64Var(&fl.c, "c", 0, "candy count")
	flag.Int64Var(&fl.m, "m", 0, "money given to the machine")
}

var fl ClientFlags

func getCert(certFile, keyfile string) (c tls.Certificate, err error) {
	if certFile != "" && keyfile != "" {
		c, err = tls.LoadX509KeyPair(certFile, keyfile)
		if err != nil {
			log.Fatalf("error loading key pair: %v\n", err)
		}
	} else {
		err = fmt.Errorf("no certificate")
	}
	return
}

func ClientCertReqFunc(certFile, keyfile string) func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
	c, err := getCert(certFile, keyfile)

	return func(certReq *tls.CertificateRequestInfo) (*tls.Certificate, error) {
		if err != nil || certFile == "" {
			log.Fatalln("no certificate", err)
		}
		return &c, nil
	}
}

func NewClient() *http.Client {
	data, err := os.ReadFile(MiniCA)
	if err != nil {
		log.Fatalln(err)
	}
	cp, err := x509.SystemCertPool()
	if err != nil {
		log.Fatalln(err)
	}

	cp.AppendCertsFromPEM(data)

	tlsCfg := &tls.Config{
		RootCAs:              cp,
		GetClientCertificate: ClientCertReqFunc(CertFile, KeyFile),
	}

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsCfg,
		},
	}
}

func main() {
	flag.Parse()
	client := NewClient()

	reqBody := fmt.Sprintf(`{"money": %d, "candyType": "%s", "candyCount": %d}`, fl.m, fl.k, fl.c)
	resp, err := client.Post("https://localhost:8080/buy_candy", "application/json", strings.NewReader(reqBody))
	if err != nil {
		log.Fatalln(err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Printf("%s", respBody)
}
