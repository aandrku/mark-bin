package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	loggerHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(loggerHandler)

	app := application{logger: logger}

	logger.Info("starting server", "addr", ":4000")

	err := http.ListenAndServe(":4000", app.routes())
	log.Fatal(err)
}
