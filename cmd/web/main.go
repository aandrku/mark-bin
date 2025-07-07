package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from mark-bin"))
}

func snippetViewGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	msg := fmt.Sprintf("Displaying snippet %s...", id)
	w.Write([]byte(msg))
}

func snippetCreateGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displaying snippet create form"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Created new snippet"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", homeGet)
	mux.HandleFunc("GET /snippet/view/{id}", snippetViewGet)
	mux.HandleFunc("GET /snippet/create", snippetCreateGet)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
