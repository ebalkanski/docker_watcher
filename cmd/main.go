package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Println("starting service")

	mux := http.NewServeMux()
	mux.HandleFunc("/", Test)

	err := http.ListenAndServe(":8080", mux)
	log.Printf("Error starting server: %v\n", err)
}

func Test(w http.ResponseWriter, r *http.Request) {
	log.Println("calling Test api")

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "Hello",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error writing response: %v\n", err)
	}
}
