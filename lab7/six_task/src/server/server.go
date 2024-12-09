package main

import (
	sixtask "PP_LABS/lab7/six_task/src/struct"
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := sixtask.NewServer()

	http.HandleFunc("/ws", server.HandleConnections)
	fs := http.FileServer(http.Dir("./lab7/six_task/index.html"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))

	// Обработка HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./lab7/six_task/index.html")
	})
	go server.HandleMessages()

	fmt.Println("WebSocket server is running on ws://localhost:8080/ws")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
