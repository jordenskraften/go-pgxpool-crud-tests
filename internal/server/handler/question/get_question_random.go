package question

import (
	"encoding/json"
	"log"
	"net/http"
	usecaseQuestion "pxgpool-crud-tests/internal/usecase/question"
)

type GetRandomQuestionHandler struct {
	usecase     *usecaseQuestion.QuestionService
	HttpHandler func(w http.ResponseWriter, r *http.Request)
}

func NewGetRandomQuestionHandler(usecase *usecaseQuestion.QuestionService) *GetRandomQuestionHandler {
	handler := &GetRandomQuestionHandler{
		usecase:     usecase,
		HttpHandler: nil,
	}
	handler.HttpHandler = handler.ServeHTTP

	return handler
}

func (gqh *GetRandomQuestionHandler) GetUrlPattern() string {
	return ""
}

func (gqh *GetRandomQuestionHandler) GetHandler() func(http.ResponseWriter, *http.Request) {
	return gqh.HttpHandler
}

// когда структура удовлетворяет интерфейсу Handler реализуя ServeHTTP
// её можно кидать в хендлеры мультиплексора
func (gqh *GetRandomQuestionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetRandomQuestionHandler")
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
