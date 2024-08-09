/*
 * Candy Server
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"

	sw "ex01server/server"
)

const (
	CertFile = "../cert/localhost/cert.pem"
	KeyFile  = "../cert/localhost/key.pem"
	MiniCA   = "../cert/minica.pem"
)

type Server struct {
	httpServer http.Server
}

func NewServer() *Server {
	data, err := ioutil.ReadFile(MiniCA)
	if err != nil {
		log.Fatalln(err)
	}

	cp, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
	}

	cp.AppendCertsFromPEM(data)

	return &Server{
		httpServer: http.Server{
			Addr:    ":8080",
			Handler: sw.NewRouter(),
		},
	}
}

func main() {
	log.Printf("Server started")
	s := NewServer()
	log.Fatal(s.httpServer.ListenAndServeTLS(CertFile, KeyFile))
}