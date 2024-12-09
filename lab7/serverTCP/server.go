package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var message string = "Server message"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
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
