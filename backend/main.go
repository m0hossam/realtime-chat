package main

import (
	"fmt"
	"net/http"

	"github.com/m0hossam/realtime-chat/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	// Print the request host
	fmt.Println(r.Host)

	// Upgrade connection to WS
	wsConn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err) // Write detailed error to response
	}

	client := &websocket.Client{
		Conn: wsConn,
		Pool: pool,
	}

	// Register the client to the connection pool
	client.Pool.RegisterCh <- client
	client.Read()
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Real-time Chat Application!")
	})

	pool := websocket.NewPool()
	go pool.Start()

	// The anonymous function is just a wrapper around serveWs to pass the pool argument
	// because http.HandleFunc only accepts functions with the signature func(http.ResponseWriter, *http.Request)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	port := "8080"
	fmt.Println("Setting up routes..")
	setupRoutes()
	fmt.Println("Starting server on :" + port)
	http.ListenAndServe(":"+port, nil)
}
