package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
)

func main() {
	certFile := "../client.crt"
	keyFile := "../client.key"
	caFile := "../ca.crt"
	serverAddr := "127.0.0.1:8080"
	message := ""

	fmt.Println("Введите сообщение:")
	fmt.Fscan(os.Stdin, &message)

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load client certificate and key: %v", err)
	}

	caCert, err := os.ReadFile(caFile)

	if err != nil {
		log.Fatalf("Failed to load CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: false,
	}

	conn, err := tls.Dial("tcp", serverAddr, tlsConfig)
	if err != nil {
		fmt.Println("qwerty")
		fmt.Println(err)
		return
	}
	defer conn.Close()
	conn.Write([]byte(message))

	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		if err.Error() != "EOF" {
			fmt.Println(err)
		}
		return
	}
	fmt.Println("Получено от сервера:", string(buff[:n]))
	//tastCA
	//client
	//127.0.0.1
}
