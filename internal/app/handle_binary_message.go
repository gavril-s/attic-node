package app

import (
	"fmt"

	"github.com/gavril-s/attic-node/internal/pkg/service/ws"
)

func (h *Handler) HandleBinaryMessage(bytes []byte) error {
	chunk, err := ws.ReadChunk(bytes)
	if err != nil {
		return fmt.Errorf("error reading chunk: %w", err)
	}

	receiver, ok := h.FileReceivers[chunk.FileID]
	if !ok {
		return fmt.Errorf("no file receiver for chunk")
	}

	err = receiver.ReceiveNewChunk(chunk)
	if err != nil {
		return fmt.Errorf("error receiving chunk: %w", err)
	}

	if receiver.CanBeFinished() {
		err = receiver.Finish()
		if err != nil {
			return fmt.Errorf("error finishing receiver: %w", err)
		}
		delete(h.FileReceivers, chunk.FileID)
	}
	return nil
}
