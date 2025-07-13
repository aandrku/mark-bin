package main

import (
	"errors"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	addr := flag.String("addr", ":4000", "Usage: -addr=:4000")
	mode := flag.String("mode", "dev", "Mode of operation: must be 'dev' or 'prod'")
	flag.Parse()

	var logger *slog.Logger
	switch *mode {
	case "dev":
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	case "prod":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	default:
		return errors.New("mode of operation must be 'dev' or 'prod'")
	}

	logger.Info("starting server", "addr", *addr)

	app := application{logger: logger}
	if err := http.ListenAndServe(*addr, app.routes()); err != nil {
		return err
	}

	return nil
}
