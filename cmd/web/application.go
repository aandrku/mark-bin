package main

import (
	"log/slog"

	"github.com/aandrku/mark-bin/pkg/models"
)

type application struct {
	userModel models.UserModelInterface
	logger    *slog.Logger
}
