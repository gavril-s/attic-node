package app

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Write(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Query().Get("path")
	if path == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, err := os.Create(path)
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, r.Body)
	if err != nil {
		log.Printf("Error writing to file: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
