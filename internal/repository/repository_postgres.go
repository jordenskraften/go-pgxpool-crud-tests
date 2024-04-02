package repository

import (
	"log/slog"
	database "pxgpool-crud-tests/internal/db"
	"pxgpool-crud-tests/internal/model"
	"pxgpool-crud-tests/internal/repository/question"
)

type ExampleRepository interface {
	GetRandomQuestion() (*model.Question, error)
}

type Repository struct {
	logger       *slog.Logger
	db           *database.DB
	QuestionRepo *question.QuestionRepositoryPostgres
}

func NewRepositoryPostgres(logger *slog.Logger, db *database.DB) *Repository {
	return &Repository{
		logger:       logger,
		db:           db,
		QuestionRepo: question.NewQuestionRepositoryPostgres(db, logger),
	}
}
