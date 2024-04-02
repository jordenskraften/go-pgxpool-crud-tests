package repository

import (
	"log/slog"
	database "pxgpool-crud-tests/internal/db"
)

type Repository struct {
	logger *slog.Logger
	db     *database.DB
}

func NewRepository(logger *slog.Logger, db *database.DB) *Repository {
	return &Repository{
		logger: logger,
		db:     db,
	}
}
