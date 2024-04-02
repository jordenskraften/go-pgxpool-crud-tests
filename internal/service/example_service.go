package service

import (
	"log/slog"
	"pxgpool-crud-tests/internal/model"
	"pxgpool-crud-tests/internal/repository"
)

type ExampleService struct {
	logger *slog.Logger
	repo   *repository.Repository
}

func NewExampleService(logger *slog.Logger, repo *repository.Repository) *ExampleService {
	service := &ExampleService{
		logger: logger,
		repo:   repo,
	}

	return service
}

func (es *ExampleService) GetRandomQuestion() (*model.Question, error) {
	// Вызываем метод репозитория для получения случайного вопроса
	quest, err := es.repo.GetRandomQuestion()
	if err != nil {
		// Обработка ошибки здесь
		es.logger.Error("Failed to get random question:", err)
		return nil, err
	}
	return quest, nil
}
