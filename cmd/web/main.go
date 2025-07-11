package main

import (
	"log"
	"net/http"
)

func main() {
	app := application{}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.homeGet)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetViewGet)
	mux.HandleFunc("GET /snippet/create", app.snippetCreateGet)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
