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
	mux.HandleFunc("GET /snippet/view/{id}", a.snippetViewGet)
	mux.HandleFunc("GET /snippet/create", a.snippetCreateGet)
	mux.HandleFunc("POST /snippet/create", a.snippetCreatePost)

	mux.HandleFunc("GET /login", a.loginGet)
	mux.HandleFunc("POST /login", a.loginPost)
	mux.HandleFunc("GET /signup", a.signupGet)
	mux.HandleFunc("POST /signup", a.signupPost)

	chain := alice.New(a.recoverPanic, a.logRequest, commonHeaders)

	return a.recoverPanic(chain.Then(mux))
}
