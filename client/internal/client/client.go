package client

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type client struct {
	conn  *websocket.Conn
	mu    sync.Mutex
	done  chan bool
	Write chan string
}

func NewClient(conn *websocket.Conn) *client {
	return &client{
		conn:  conn,
		done:  make(chan bool),
		Write: make(chan string),
	}
}

func (client *client) Start() {
	go client.Read()
	// Tells the master that we are a slave
	client.conn.WriteMessage(websocket.TextMessage, []byte("slave"))
	for {
		select {
		case <-client.done:
			return
		case message := <-client.Write:
			log.Println("Writing ", message)
			client.mu.Lock()
			client.conn.WriteMessage(websocket.TextMessage, []byte(message))
			client.mu.Unlock()
		}
	}
}

func (client *client) Read() {
	for {
		_, p, _ := client.conn.ReadMessage()
		message := string(p)
		log.Println("Received", message)
	}
}
