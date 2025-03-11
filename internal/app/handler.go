package app

import (
	"fmt"

	"github.com/gavril-s/attic-node/internal/config"
	"github.com/gavril-s/attic-node/internal/pkg/service/file"
	"github.com/gorilla/websocket"
)

type Handler struct {
	Cfg  *config.Config
	Conn *websocket.Conn

	FileReceivers map[string]*file.FileReceiver
	FileSenders   map[string]*file.FileSender
}

func NewHandler(cfg *config.Config, conn *websocket.Conn) *Handler {
	return &Handler{
		Cfg:  cfg,
		Conn: conn,
	}
}

func (h *Handler) HandleMessage(messageType int, bytes []byte) error {
	switch messageType {
	case websocket.BinaryMessage:
		return h.HandleBinaryMessage(bytes)
	case websocket.TextMessage:
		return h.HandleTextMessage(bytes)
	default:
		return fmt.Errorf("unsupported message type")
	}
}
