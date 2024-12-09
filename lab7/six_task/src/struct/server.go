package sixtask

import (
	observer "PP_LABS/lab7/six_task/src/interface"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	observers   []observer.Observer
	LastMessage Message
	mu          sync.Mutex
	upgrader    websocket.Upgrader
	broadcast   chan Message
}

func NewServer() *Server {
	return &Server{
		observers:   []observer.Observer{},
		LastMessage: Message{""},
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		broadcast: make(chan Message),
	}
}

func (server *Server) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Error upgrading to WebSocket: %v", err)
		return
	}
	defer conn.Close()
	client := NewClient(server, conn)
	server.RegisterObserver(client)
	fmt.Println("New client connected")

	for {
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			server.RemoveObserver(client)
			break
		}
		server.broadcast <- message
	}
}

// Метод для рассылки сообщений всем клиентам
func (server *Server) HandleMessages() {
	for {
		message := <-server.broadcast
		server.mu.Lock()
		server.SetLastMessage(message)
		server.NotifyObservers()
		server.mu.Unlock()
	}
}

func (s *Server) RegisterObserver(o observer.Observer) {
	s.observers = append(s.observers, o)
}

func (s *Server) RemoveObserver(o observer.Observer) {
	for i, obs := range s.observers {
		if obs == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Server) NotifyObservers() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Server) GetLastMessage() Message {
	return s.LastMessage
}

func (s *Server) SetLastMessage(message Message) {
	s.LastMessage = message
}

func (s *Server) Test(o *observer.Observer) {
	fmt.Println("fff")
}
