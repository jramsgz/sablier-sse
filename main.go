package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// Set necessary headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Flush the headers
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	for {
		// Send the data in the format "data: message"
		fmt.Fprintf(w, "data: It works\n\n")
		flusher.Flush() // Send it to the client

		// Wait 2 seconds before sending the next message
		time.Sleep(2 * time.Second)
	}
}

func main() {
	http.HandleFunc("/sse/events", sseHandler)

	fmt.Println("Starting SSE server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
