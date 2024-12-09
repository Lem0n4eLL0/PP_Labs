package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	message := ""
	fmt.Println("Введите сообщение:")
	fmt.Fscan(os.Stdin, &message)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
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

}
