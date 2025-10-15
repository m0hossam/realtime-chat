package main

import (
	"fmt"
	"net/http"

	"github.com/m0hossam/realtime-chat/pkg/websocket"
)

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Real-time Chat Application!")
	})

	pool := websocket.NewPool()
	go pool.Start()

	// The anonymous function is just a wrapper around serveWs to pass the pool argument
	// because http.HandleFunc only accepts functions with the signature func(http.ResponseWriter, *http.Request)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(pool, w, r)
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
