package websocket

import (
	"github.com/MikeDevresse/websocket-scale/ws/internal/router"
	"github.com/gorilla/websocket"
	"sync"
)

type client struct {
	identifier string
	conn       *websocket.Conn
	mu         sync.Mutex
}

func NewClient(identifier string, conn *websocket.Conn) *client {
	return &client{
		identifier: identifier,
		conn:       conn,
	}
}

// Write sends a message through the websocket connection
func (client *client) Write(message string) {
	// We lock the mutex in order to prevent concurrency
	client.mu.Lock()
	client.conn.WriteMessage(websocket.TextMessage, []byte(message))
	client.mu.Unlock()
}

// Read gets all the messages that the client send
func (client *client) Read() {
	// On quit, we close the connection
	defer func() {
		client.conn.Close()
	}()

	for {
		_, p, err := client.conn.ReadMessage()
		// If an error occurred then it means that the connection is closed, we close the function
		if err != nil {
			return
		}
		router.Handle(string(p))
	}
}
