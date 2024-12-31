package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response struct for JSON response
type Response struct {
	Message string `json:"message"`
}

// HelloHandler handles the /hello endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello World!"}
	json.NewEncoder(w).Encode(response)
}

// main function to start the server
func main() {
	http.HandleFunc("/hello", HelloHandler)

	// Start the server on port 8080
	port := "8080"
	log.Printf("Server starting at port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}