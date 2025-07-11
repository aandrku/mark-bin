package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "Usage: -addr=:4000")
	mode := flag.String("mode", "dev", "Mode of operation: must be 'dev' or 'prod'")
	flag.Parse()

	_ = mode

	loggerHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(loggerHandler)

	app := application{logger: logger}

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	log.Fatal(err)
}
