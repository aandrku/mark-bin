package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homeGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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
