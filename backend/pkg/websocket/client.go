package websocket

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn   *websocket.Conn
	Pool   *Pool
	SendCh chan string
}

func (c *Client) Read() {
	// Ensure client is unregistered and connection is closed
	defer func() {
		c.Pool.UnregisterCh <- c
		c.Conn.Close()
	}()

	// Continuously read incoming messages, and broadcast them to all clients
	for {
		_, b, err := c.Conn.ReadMessage()
		if err != nil {
			// Only log abnormal errors, because normal closure and going away also return an error
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				fmt.Println("read error: ", err)
			}
			break
		}
		c.Pool.BroadcastCh <- string(b)
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close() // This is idempotent, so doing it in both Read() and Write() is safe
	}()

	// Write messages received from the channel until channel is closed
	for msg := range c.SendCh {
		if err := c.Conn.WriteMessage(1, []byte(msg)); err != nil {
			break
		}
	}

	c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
}
