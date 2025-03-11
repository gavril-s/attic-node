package main

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gavril-s/attic-master/api"
	"github.com/gavril-s/attic-node/internal/config"
	"github.com/gorilla/websocket"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	u := url.URL{Scheme: "ws", Host: cfg.MasterURL, Path: "/ws"}
	log.Printf("Connecting to master server at %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error from master:", err)
				return
			}

			if messageType == websocket.TextMessage {
				var textMessage api.Message
				err = json.Unmarshal(message, &textMessage)
				if err != nil {
					return
				}
				switch textMessage.MessageType {

				}
			}

			log.Printf("Received from master: %s", message)
		}
	}()

	err = conn.WriteMessage(websocket.BinaryMessage, []byte("Node connected and active"))
	if err != nil {
		log.Println("Initial write error:", err)
		return
	}

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("Interrupt received, closing connection...")

			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Write close error:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
