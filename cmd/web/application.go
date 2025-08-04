package main

import (
	"log/slog"

	"github.com/aandrku/mark-bin/pkg/models"
)

type application struct {
	snippetModel *models.SnippetModel
	userModel    models.UserModelInterface
	logger       *slog.Logger
}
