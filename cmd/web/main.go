package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/aandrku/mark-bin/pkg/models"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	addr := flag.String("addr", ":4000", "Usage: -addr=:4000")
	mode := flag.String("mode", "dev", "Mode of operation: must be 'dev' or 'prod'")
	dsn := flag.String("dsn", "", "Usage: -dsn=postgres://<username>:<password>@<ip>/<db-name>")
	flag.Parse()

	var logger *slog.Logger
	switch *mode {
	case "dev":
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	case "prod":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	default:
		return fmt.Errorf("mode of operation must be 'dev' or 'prod', not %s", *mode)
	}

	db, err := openDB(*dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	_ = db

	logger.Info("starting server", "addr", *addr)

	um := &models.UserModel{DB: db}
	app := application{
		userModel: um,
		logger:    logger}
	if err := http.ListenAndServe(*addr, app.routes()); err != nil {
		return err
	}

	return nil
}
