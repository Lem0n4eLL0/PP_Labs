package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var message string = "Server message"

func main() {
	certFile := "../server.crt"
	keyFile := "../server.key"
	caFile := "../ca.crt"
	serverAddr := "127.0.0.1:8080"

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load server certificate and key: %v", err)
	}

	caCert, err := os.ReadFile(caFile)
	if err != nil {
		log.Fatalf("Failed to load CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	listener, err := tls.Listen("tcp", serverAddr, tlsConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	connections := make(chan net.Conn)
	defer close(connections)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case conn := <-connections:
				handleConnection(conn)
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Println(err)
					continue
				}
			}
			fmt.Println(conn.LocalAddr().String())
			select {
			case <-ctx.Done():
				conn.Close()
			case connections <- conn:
			}
		}
	}()
	<-ctx.Done()
	wg.Wait()
	fmt.Println("Server stopped")
}

func handleConnection(con net.Conn) {
	defer con.Close()
	buff := make([]byte, 1024)
	con.Read(buff)
	fmt.Println(string(buff[:]))
	_, err := con.Write([]byte(message))
	if err != nil {
		fmt.Println(err)
	}
}
