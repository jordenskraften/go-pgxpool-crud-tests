package question

import (
	"encoding/json"
	"net/http"
	usecaseQuestion "pxgpool-crud-tests/internal/usecase/question"
)

type GetRandomQuestionHandler struct {
	PatternUrl  string
	usecase     *usecaseQuestion.QuestionService
	HttpHandler func(w http.ResponseWriter, r *http.Request)
}

func NewGetRandomQuestionHandler(patternUrl string, usecase *usecaseQuestion.QuestionService) *GetRandomQuestionHandler {
	handler := &GetRandomQuestionHandler{
		PatternUrl:  patternUrl,
		usecase:     usecase,
		HttpHandler: nil,
	}
	handler.HttpHandler = handler.ServeHTTP

	return handler
}

// когда структура удовлетворяет интерфейсу Handler реализуя ServeHTTP
// её можно кидать в хендлеры мультиплексора
func (gqh *GetRandomQuestionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	quest, err := gqh.usecase.GetRandomQuestion()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := ErrorResponse{Error: err.Error()}
		errorJSON, _ := json.Marshal(errorResponse) // Ошибку можно не обрабатывать для логирования
		w.Write(errorJSON)                          // Отправляем тот же JSON клиенту
		return
	}

	response := QuestionResponse{
		Id:            quest.Id,
		TopicId:       quest.Topic_id,
		QuestionTitle: quest.Question_title,
		QuestionText:  quest.Question_text,
		AnswerText:    quest.Answer_text,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		// Обработка ошибки сериализации, если требуется
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON) // Отправляем сериализованный JSON клиенту
}

func (gqh *GetRandomQuestionHandler) GetUrlPattern() string {
	return gqh.PatternUrl
}

func (gqh *GetRandomQuestionHandler) GetHandler() func(http.ResponseWriter, *http.Request) {
	return gqh.HttpHandler
}
