package websocket

import "fmt"

type Pool struct {
	RegisterCh   chan *Client
	UnregisterCh chan *Client
	BroadcastCh  chan string
	Clients      map[*Client]bool
}

func NewPool() *Pool {
	return &Pool{
		// All channels are unbuffered to block and allow only one client/message at a time (acting like a mutex)
		RegisterCh:   make(chan *Client),
		UnregisterCh: make(chan *Client),
		BroadcastCh:  make(chan string),
		Clients:      make(map[*Client]bool),
	}
}

func (pool *Pool) Start() {
	for { // Continuously listening for any activity on the channels
		select {
		case client := <-pool.RegisterCh: // New client registered
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
		case client := <-pool.UnregisterCh: // Client unregistered
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
		case msg := <-pool.BroadcastCh: // Received new message that needs to be broadcasted to all clients
			for c := range pool.Clients {
				select {
				case c.SendCh <- msg: // Asynchronously enqueue the message to each client's send-channel
				default: // If a client is blocking, close his buffered channel and drop him
					close(c.SendCh)
					delete(pool.Clients, c)
				}
			}
		}
	}
}
