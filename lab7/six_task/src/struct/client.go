package sixtask

import (
	observer "PP_LABS/lab7/six_task/src/interface"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	observer    observer.Observable
	LastMessage Message
	conn        *websocket.Conn
}

func (c *Client) Update(o observer.Observable) {

	server, ok := o.(*Server)
	if !ok {
		fmt.Println("Невозможно преобразовать Observable в Server")
		return
	}
	c.LastMessage = server.GetLastMessage()
	messageJSON, err := json.Marshal(c.LastMessage)
	if err != nil {
		fmt.Printf("Error marshalling message: %v\n", err)
		return
	}

	err = c.conn.WriteMessage(websocket.TextMessage, messageJSON)
	if err != nil {
		fmt.Printf("Error sending message to client: %v\n", err)
	}
}

func NewClient(o observer.Observable, c *websocket.Conn) *Client {
	client := &Client{
		observer:    o,
		LastMessage: Message{""},
		conn:        c,
	}
	return client
}
