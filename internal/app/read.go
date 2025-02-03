package app

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gavril-s/attic-node/internal/config"
)

func Read(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Query().Get("path")
	if path == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	flusher, ok := w.(http.Flusher)
	if !ok {
		log.Printf("Error creating http.Flusher: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	buf := make([]byte, config.FileChunkSize)
	reader := bufio.NewReader(file)
	for {
		bytesRead, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading file: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			break
		}

		if bytesRead == len(buf) {
			io.Writer.Write(w, buf)
		} else {
			io.Writer.Write(w, buf[:bytesRead])
		}
		flusher.Flush()
	}
}
