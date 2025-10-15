package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { // allowing connections from any origin
		return true
	},
}

func ServeWs(pool *Pool, w http.ResponseWriter, r *http.Request) {
	// Upgrade connection to WS
	wsConn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err) // Write detailed error to response
	}

	// Create and register the client to the connection pool
	client := &Client{
		Conn:   wsConn,
		Pool:   pool,
		SendCh: make(chan string, 256),
	}
	client.Pool.RegisterCh <- client

	go client.Read()
	go client.Write()
}
