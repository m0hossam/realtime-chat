package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

func (c *Client) Read() {
	// Ensure client is unregistered and connection is closed when Read() returns
	defer func() {
		c.Pool.UnregisterCh <- c
		c.Conn.Close()
	}()

	// Continuously read messages coming from the frontend, and broadcast them to all clients
	for {
		_, b, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		c.Pool.BroadcastCh <- string(b)
		fmt.Printf("Message Received: %+v\n", string(b))
	}
}
