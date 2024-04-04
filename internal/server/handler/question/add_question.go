package question

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pxgpool-crud-tests/internal/model"
	usecaseQuestion "pxgpool-crud-tests/internal/usecase/question"
)

type AddQuestionHandler struct {
	PatternUrl  string
	usecase     *usecaseQuestion.QuestionService
	HttpHandler func(w http.ResponseWriter, r *http.Request)
}

func NewAddQuestionHandler(patternUrl string, usecase *usecaseQuestion.QuestionService) *AddQuestionHandler {
	handler := &AddQuestionHandler{
		PatternUrl:  patternUrl,
		usecase:     usecase,
		HttpHandler: nil,
	}
	handler.HttpHandler = handler.ServeHTTP
	return handler
}

func (aqh *AddQuestionHandler) GetUrlPattern() string {
	return aqh.PatternUrl
}

func (aqh *AddQuestionHandler) GetHandler() func(http.ResponseWriter, *http.Request) {
	return aqh.HttpHandler
}

func (aqh *AddQuestionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Парсинг JSON из тела запроса
	var questionDTO model.QuestionDTO
	if err := json.NewDecoder(r.Body).Decode(&questionDTO); err != nil {
		http.Error(w, "Ошибка при чтении JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Преобразование QuestionDTO в model.Question
	questionModel := &model.QuestionDTO{
		TopicID:       questionDTO.TopicID,
		QuestionTitle: questionDTO.QuestionTitle,
		QuestionText:  questionDTO.QuestionText,
		AnswerText:    questionDTO.AnswerText,
	}

	// Добавление вопроса через usecase
	id, err := aqh.usecase.AddNewQeustion(questionModel)
	if err != nil {
		http.Error(w, "Ошибка при добавлении вопроса: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Формирование JSON-ответа
	type Response struct {
		Message string `json:"message"`
	}
	response := Response{
		Message: fmt.Sprintf("Вопрос успешно добавлен, ID вопроса: %d", id),
	}

	// Кодирование JSON-ответа
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Ошибка при кодировании JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка JSON-ответа клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
