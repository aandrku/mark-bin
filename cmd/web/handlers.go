package main

import (
	"fmt"
	"net/http"

	"github.com/aandrku/mark-bin/ui/templ/pages"
)

func homeGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	page := pages.Home()

	render(w, page)
}

func snippetViewGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	fmt.Fprintf(w, "Displaying snippet %s", id)
}

func snippetCreateGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displaying snippet create form"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Created new snippet"))
}
