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
		case message := <-pool.BroadcastCh: // Received new message that needs to be broadcasted to all clients
			for c := range pool.Clients {
				if err := c.Conn.WriteMessage(1, []byte(message)); err != nil {
					fmt.Println(err)
					return // Terminate all connections if one fails (overkill)
				}
			}
		}
	}
}
