package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/RickardA/stop-watch/internal/app/gui"
)

func main() {
	port := "8081"

	stopChan := make(chan int, 1)

	// Setup Client
	srv := http.NewServeMux()

	srv.Handle("/query", stopHandler{StopChan: &stopChan})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	go http.ListenAndServe(":"+port, srv)

	gui.NewApp("Stop Watch", &stopChan)
}

type stopHandler struct {
	StopChan *chan int
}

func (sh stopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data stopData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(data.Lane)

	if data.Lane == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	*sh.StopChan <- data.Lane
}

type stopData struct {
	Lane int `json:"lane"`
}
