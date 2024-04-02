package example

import (
	"context"
	"log/slog"
	database "pxgpool-crud-tests/internal/db"
	"pxgpool-crud-tests/internal/model"

	"github.com/Masterminds/squirrel"
)

type ExampleRepositoryPostgres struct {
	logger *slog.Logger
	db     *database.DB
}

func NewExampleRepositoryPostgres(db *database.DB, logger *slog.Logger) *ExampleRepositoryPostgres {
	return &ExampleRepositoryPostgres{
		db:     db,
		logger: logger,
	}
}

func (ex *ExampleRepositoryPostgres) GetRandomQuestion() (*model.Question, error) {
	// Используем Squirrel для построения запроса
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.
		Select("id", "question_title", "question_text", "answer_text").
		From("Question").
		OrderBy("RANDOM()").
		Limit(1).
		ToSql()

	if err != nil {
		ex.logger.Error("Failed to build query: ", err)
		return nil, err
	}

	// Подготовка структуры для хранения результата
	var question model.Question

	// Выполнение запроса
	err = ex.db.Pool.QueryRow(context.Background(), query, args...).Scan(
		&question.Id,
		&question.Question_title,
		&question.Question_text,
		&question.Answer_text,
	)

	ex.logger.Info("Fetched question: %+v\n", question)

	if err != nil {
		ex.logger.Error("Failed to fetch random question: ", err)
		return nil, err
	}

	return &question, nil
}
