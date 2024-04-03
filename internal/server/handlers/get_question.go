package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	usecaseQuestion "pxgpool-crud-tests/internal/usecase/question"
)

type GetQuestionHandler struct {
	router  *http.ServeMux
	usecase *usecaseQuestion.QuestionService
	logger  *slog.Logger
}

// QuestionResponse структура для отправки вопроса клиенту.
type QuestionResponse struct {
	Id            int    `json:"id"`
	TopicId       int    `json:"topic_id"`
	QuestionTitle string `json:"question_title"`
	QuestionText  string `json:"question_text"`
	AnswerText    string `json:"answer_text"`
}

// ErrorResponse структура для отправки ошибки клиенту.
type ErrorResponse struct {
	Error string `json:"error"`
}

func NewGetQuestionHandler(router *http.ServeMux, usecase *usecaseQuestion.QuestionService, logger *slog.Logger) *GetQuestionHandler {
	return &GetQuestionHandler{
		router:  router,
		usecase: usecase,
		logger:  logger,
	}
}

// когда структура удовлетворяет интерфейсу Handler реализуя ServeHTTP
// её можно кидать в хендлеры мультиплексора
func (gqh *GetQuestionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	quest, err := gqh.usecase.GetRandomQuestion()
	if err != nil {
		gqh.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := ErrorResponse{Error: err.Error()}
		errorJSON, _ := json.Marshal(errorResponse) // Ошибку можно не обрабатывать для логирования
		gqh.logger.Info(string(errorJSON))          // Логируем JSON с ошибкой
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

	gqh.logger.Info("Sending response: ", slog.Int("id", response.Id), slog.String("title", response.QuestionTitle), slog.String("text", response.QuestionText)) // Логируем JSON перед отправкой
	responseJSON, err := json.Marshal(response)
	if err != nil {
		// Обработка ошибки сериализации, если требуется
		gqh.logger.Error("Failed to serialize response: ", err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON) // Отправляем сериализованный JSON клиенту
}
