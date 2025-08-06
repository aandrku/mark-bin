package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aandrku/mark-bin/ui/templ/pages"
)

func (a *application) homeGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	snippets, err := a.snippetModel.Latest(ctx, 10)
	if err != nil {
		a.serverError(w, r, err)
	}

	page := pages.Home(snippets)

	if err := render(w, page); err != nil {
		a.serverError(w, r, err)
	}
}

func (a *application) snippetsViewsGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ctx := r.Context()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		a.logger.Error("Bad request")
		return
	}

	s, err := a.snippetModel.GetWithUsername(ctx, idInt)
	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		a.logger.Error("Bad request")
		return
	}

	page := pages.SnippetView(s)
	render(w, page)

}

func (a *application) snippetsCreateGet(w http.ResponseWriter, r *http.Request) {
	page := pages.SnippetCreate()

	render(w, page)
}

func (a *application) snippetsPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Created new snippet"))
}

// loginGet shows login form to the user.
func (a *application) loginGet(w http.ResponseWriter, r *http.Request) {
	page := pages.Login()

	render(w, page)
}

// loginPost logs the user in.
func (a *application) loginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Log user in \n")

}

// signupGet shows sign up form to the user.
func (a *application) signupGet(w http.ResponseWriter, r *http.Request) {
	page := pages.Signup()

	render(w, page)
}

// signupPost registers new user.
func (a *application) signupPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register new user")
}
