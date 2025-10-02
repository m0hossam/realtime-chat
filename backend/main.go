package main

import (
	"fmt"
	"log"
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

func reader(conn *websocket.Conn) {
	for {
		// Read message
		messageType, b, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// Print the message to the console
		fmt.Println(string(b))

		// Write same message back
		if err := conn.WriteMessage(messageType, b); err != nil {
			log.Println(err)
			return
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	// Print the request host
	fmt.Println(r.Host)

	// Upgrade connection to WS
	ws, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// Listen to messages and write them back
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Real-time Chat Application!")
	})

	http.HandleFunc("/ws", serveWs)
}

func main() {
	port := "8080"
	fmt.Println("Setting up routes..")
	setupRoutes()
	fmt.Println("Starting server on :" + port)
	http.ListenAndServe(":"+port, nil)
}
