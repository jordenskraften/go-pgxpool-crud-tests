package repository

import (
	"log/slog"
	database "pxgpool-crud-tests/internal/db"
	"pxgpool-crud-tests/internal/model"
	"pxgpool-crud-tests/internal/repository/example"
)

type ExampleRepository interface {
	GetRandomQuestion() (*model.Question, error)
}

type Repository struct {
	logger      *slog.Logger
	db          *database.DB
	ExampleRepo *example.ExampleRepositoryPostgres
}

func NewRepositoryPostgres(logger *slog.Logger, db *database.DB) *Repository {
	return &Repository{
		logger:      logger,
		db:          db,
		ExampleRepo: example.NewExampleRepositoryPostgres(db, logger),
	}
}
