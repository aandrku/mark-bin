package main

import (
	"log"
	"net/http"
)

func homeGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from mark-bin"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", homeGet)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
