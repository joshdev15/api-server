package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Good(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]interface{}{"title": "Simple API Rest Go example", "success": "Hola server en Golang!", "statusCode": http.StatusOK})
	if err != nil {
		panic(err)
	}
}

func Bad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(w).Encode(map[string]interface{}{"title": "Simple API Rest Go example", "error": "Esto es un error!", "statusCode": http.StatusInternalServerError})
	if err != nil {
		panic(err)
	}
}

func main() {
	port := ":8080"

	/* Router */
	mux := http.NewServeMux()

	/* Good Response */
	mux.HandleFunc("/", Good)

	/* Bad Response */
	mux.HandleFunc("/error", Bad)

	fmt.Printf("Your server is running! - http://localhost%s\n", port)

	/* Run Server */
	log.Fatal(http.ListenAndServe(port, mux))
}

/* Made in Go 1.19 */
