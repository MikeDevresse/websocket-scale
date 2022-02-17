package main

import (
	"github.com/MikeDevresse/websocket-scale/client/internal/client"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctrlc := make(chan os.Signal)
	signal.Notify(ctrlc, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ctrlc
		os.Exit(1)
	}()

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:49163", Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Could not connect to master:", err)
	}

	wsClient := client.NewClient(conn)
	wsClient.Start()
}
