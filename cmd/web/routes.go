package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (a *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", a.homeGet)
	mux.HandleFunc("GET /snippets/views/{id}", a.snippetsViewsGet)
	mux.HandleFunc("GET /snippets/create", a.snippetsCreateGet)
	mux.HandleFunc("POST /snippets", a.snippetsPost)

	mux.HandleFunc("GET /login", a.loginGet)
	mux.HandleFunc("POST /login", a.loginPost)
	mux.HandleFunc("GET /signup", a.signupGet)
	mux.HandleFunc("POST /signup", a.signupPost)

	chain := alice.New(authenticate, a.recoverPanic, a.logRequest, commonHeaders)

	return a.recoverPanic(chain.Then(mux))
}
