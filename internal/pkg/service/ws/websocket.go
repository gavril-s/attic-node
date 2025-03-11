package ws

import (
	"encoding/json"

	"github.com/gavril-s/attic-master/api"
	"github.com/gorilla/websocket"
)

func SendMessage(conn *websocket.Conn, message api.Message) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return conn.WriteMessage(websocket.TextMessage, bytes)
}

func SendChunk(conn *websocket.Conn, message api.Chunk) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return conn.WriteMessage(websocket.BinaryMessage, bytes)
}

func ReadMessage(bytes []byte) (api.Message, error) {
	var message api.Message
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		return api.Message{}, err
	}
	return message, nil
}

func ReadChunk(bytes []byte) (api.Chunk, error) {
	var message api.Chunk
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		return api.Chunk{}, err
	}
	return message, nil
}
