package repository

import (
	"context"
	"pxgpool-crud-tests/internal/model"
)

type ExampleRepository interface {
	GetRandomQuestion() (*model.Question, error)
}

func (repo *Repository) GetRandomQuestion() (*model.Question, error) {
	// Запрос для получения случайной записи
	query := "SELECT id, question_title, question_text, answer_text FROM Question ORDER BY RANDOM() LIMIT 1"

	// Подготовка структуры для хранения результата
	var question model.Question

	// Выполнение запроса
	err := repo.db.Pool.QueryRow(context.Background(), query).Scan(
		&question.Id,
		&question.Question_title,
		&question.Question_text,
		&question.Answer_text,
	)
	repo.logger.Info("Fetched question: %+v\n", question)
	if err != nil {
		repo.logger.Error("Failed to fetch random question: ", err)
		return nil, err
	}

	return &question, nil
}
