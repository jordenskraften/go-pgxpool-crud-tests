package question

import (
	"encoding/json"
	"log"
	"net/http"
	usecaseQuestion "pxgpool-crud-tests/internal/usecase/question"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type GetQuestionByIdHandler struct {
	usecase     *usecaseQuestion.QuestionService
	HttpHandler func(w http.ResponseWriter, r *http.Request)
}

func NewGetQuestionByIdHandler(usecase *usecaseQuestion.QuestionService) *GetQuestionByIdHandler {
	handler := &GetQuestionByIdHandler{
		usecase:     usecase,
		HttpHandler: nil,
	}
	handler.HttpHandler = handler.ServeHTTP
	return handler
}

func (gqh *GetQuestionByIdHandler) GetUrlPattern() string {
	return ""
}

func (gqh *GetQuestionByIdHandler) GetHandler() func(http.ResponseWriter, *http.Request) {
	return gqh.HttpHandler
}

func (gqh *GetQuestionByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//idQuery := r.PathValue("id")

	idQuery := chi.URLParam(r, "id") //chi.URLParam(r, "id") //
	log.Printf(idQuery)
	// Преобразование строки в int
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		// Обработка ошибки сериализации, если требуется
		http.Error(w, "Invalid question id", http.StatusInternalServerError)
		return
	}

	quest, err := gqh.usecase.GetQuestionById(id)
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
