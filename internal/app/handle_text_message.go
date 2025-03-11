package app

import (
	"fmt"

	"github.com/gavril-s/attic-master/api"
	"github.com/gavril-s/attic-node/internal/pkg/service/file"
	"github.com/gavril-s/attic-node/internal/pkg/service/ws"
)

func (h *Handler) HandleTextMessage(bytes []byte) error {
	message, err := ws.ReadMessage(bytes)
	if err != nil {
		return fmt.Errorf("error reading message: %w", err)
	}

	switch message.MessageType {
	case api.MessageTypeAcknowledgement:
		return h.handleAcknowledgement(message)
	case api.MessageTypeFileRequest:
		return h.handleFileRequest(message)
	case api.MessageTypeMasterFileHeader:
		return h.handleMasterFileHeader(message)
	default:
		return fmt.Errorf("unsupported message type: %v", message.MessageType)
	}
}

func (h *Handler) handleAcknowledgement(_ api.Message) error {
	return nil
}

func (h *Handler) handleFileRequest(_ api.Message) error {
	return nil
}

func (h *Handler) handleMasterFileHeader(message api.Message) error {
	header := message.MasterFileHeader
	if _, ok := h.FileReceivers[header.FileID]; ok {
		return fmt.Errorf("file receiver for file already exists")
	}

	var err error
	h.FileReceivers[header.FileID], err = file.NewFileReceiver(h.Cfg, header)
	if err != nil {
		return fmt.Errorf("error creating file receiver: %w", err)
	}
	return nil
}
