package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", homeGet)
	mux.HandleFunc("GET /snippet/view/{id}", snippetViewGet)
	mux.HandleFunc("GET /snippet/create", snippetCreateGet)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
