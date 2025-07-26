package main

import (
	"fmt"
	"net/http"
	"strconv"

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
	ctx := r.Context()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting this snipet %v", err)
	}

	s, err := a.snippetModel.GetWithUsername(ctx, idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting this snipet %v", err)
	}

	page := pages.SnippetView(s)
	render(w, page)

}

func (a *application) snippetCreateGet(w http.ResponseWriter, r *http.Request) {
	page := pages.SnippetCreate()

	render(w, page)
}

func (a *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Created new snippet"))
}
