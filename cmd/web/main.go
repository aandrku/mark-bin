package main

import (
	"log"
	"net/http"
)

func main() {
	app := application{}

	err := http.ListenAndServe(":4000", app.routes())
	log.Fatal(err)
}
