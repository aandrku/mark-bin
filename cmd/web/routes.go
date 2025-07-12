package main

import "net/http"

func (a *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", a.homeGet)
	mux.HandleFunc("GET /snippet/view/{id}", a.snippetViewGet)
	mux.HandleFunc("GET /snippet/create", a.snippetCreateGet)
	mux.HandleFunc("POST /snippet/create", a.snippetCreatePost)

	return commonHeaders(mux)
}
