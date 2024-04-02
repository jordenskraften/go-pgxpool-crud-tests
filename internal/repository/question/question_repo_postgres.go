package question

import (
	"context"
	database "pxgpool-crud-tests/internal/db"
	"pxgpool-crud-tests/internal/model"

	"github.com/Masterminds/squirrel"
)

type QuestionRepositoryPostgres struct {
	db *database.DB
}

func NewQuestionRepositoryPostgres(db *database.DB) *QuestionRepositoryPostgres {
	return &QuestionRepositoryPostgres{
		db: db,
	}
}

func (ex *QuestionRepositoryPostgres) GetRandomQuestion() (*model.Question, error) {
	// Используем Squirrel для построения запроса
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.
		Select("id", "question_title", "question_text", "answer_text").
		From("Question").
		OrderBy("RANDOM()").
		Limit(1).
		ToSql()

	if err != nil {
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

	if err != nil {
		return nil, err
	}

	return &question, nil
}
