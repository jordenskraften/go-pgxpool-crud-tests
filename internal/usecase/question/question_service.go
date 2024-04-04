package question

import (
	"pxgpool-crud-tests/internal/model"
)

type QuestionRepository interface {
	GetRandomQuestion() (*model.Question, error)
	GetQuestionById(id int) (*model.Question, error)
	AddQuestion(*model.QuestionDTO) (int, error)
}
type QuestionService struct {
	QuestionRepository
}

func NewQuestionService(repo QuestionRepository) *QuestionService {
	service := &QuestionService{
		QuestionRepository: repo,
	}

	return service
}

func (es *QuestionService) GetRandomQuestion() (*model.Question, error) {
	// Вызываем метод репозитория для получения случайного вопроса
	quest, err := es.QuestionRepository.GetRandomQuestion()
	if err != nil {
		// Обработка ошибки здесь
		return nil, err
	}
	return quest, nil
}

func (es *QuestionService) GetQuestionById(id int) (*model.Question, error) {
	// Вызываем метод репозитория для получения случайного вопроса
	quest, err := es.QuestionRepository.GetQuestionById(id)
	if err != nil {
		// Обработка ошибки здесь
		return nil, err
	}
	return quest, nil
}

func (es *QuestionService) AddNewQeustion(questionDTO *model.QuestionDTO) (int, error) {
	// Вызываем метод репозитория для получения случайного вопроса
	id, err := es.QuestionRepository.AddQuestion(questionDTO)
	if err != nil {
		// Обработка ошибки здесь
		return -1, err
	}
	return id, nil
}
