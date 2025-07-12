package main

import (
	"fmt"
	"net/http"

	"github.com/aandrku/mark-bin/ui/templ/pages"
)

func (a *application) homeGet(w http.ResponseWriter, r *http.Request) {
	page := pages.Home()

	if err := render(w, page); err != nil {
		a.serverError(w, r, err)
	}
}

func (a *application) snippetViewGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	fmt.Fprintf(w, "Displaying snippet %s", id)
}

func (a *application) snippetCreateGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displaying snippet create form"))
}

func (a *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Created new snippet"))
}
