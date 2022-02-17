package main

import (
	"fmt"
	"github.com/MikeDevresse/websocket-scale/ws/internal/websocket"
	"log"
	"net/http"
)

func main() {
	initWebsocket()

	port := 8000
	log.Printf("Running on http://127.0.0.1:%v", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal("An error occurred while trying to start the webserver: \"", err, "\"")
	}
}

func initWebsocket() {
	server := websocket.NewServer()
	count := 0

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r)
		if err != nil {
			fmt.Fprintf(w, "%v\n", err)
			return
		}
		log.Println("New client")

		client := websocket.NewClient("a", conn)
		server.AddClient(client)
		count = count + 1
		client.Write(fmt.Sprintf("Server %d", count))
	})
}
