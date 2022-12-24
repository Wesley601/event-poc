package main

import (
	"fmt"
	"log"
	"net/http"
	"wesley601/event-driver/internal/sse"
	"wesley601/event-driver/pkg/bus"
)

type Percent struct {
	ProgressPercentage int `json:"progressPercentage"`
}

func main() {
	eBus, err := bus.New()
	if err != nil {
		panic(err)
	}
	defer eBus.Close()

	// get broker
	bk := sse.New()

	http.HandleFunc("/progress", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("connected!")
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		// Each connection registers its own message channel with the broker's connections registry
		messageChan := make(chan []byte)

		// Signal the broker that we have a new Connection
		bk.NewClients <- messageChan

		// Remove this client from the map of connected clients
		// when this handler exits.
		defer func() {
			bk.ClosingClients <- messageChan
		}()

		ch, err := eBus.Subscribe("user.created")
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		flusher.Flush()

		for msg := range ch {
			fmt.Println("message recived")
			w.Write([]byte("event: userCreated\n"))
			w.Write([]byte(fmt.Sprintf("data: %s\n", <-messageChan)))
			w.Write([]byte("\n"))
			flusher.Flush()
		}

		w.Write([]byte("event: done\n"))
		w.Write([]byte("data: {}\n"))
		w.Write([]byte("\n"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
