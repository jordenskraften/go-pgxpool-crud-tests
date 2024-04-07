package server

import (
	"net/http"
	"pxgpool-crud-tests/internal/server/handler/question"
	"pxgpool-crud-tests/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func NewRouter(usecase *usecase.Usecase) chi.Router {
	r := chi.NewRouter()
	InitRouter(usecase, r)
	return r
}

func InitRouter(usecase *usecase.Usecase, router chi.Router) {
	//тут нарн буду инициализировать роуты все
	questionRandomHandler := question.NewGetRandomQuestionHandler(&usecase.QuestionService)
	questionByIdHandler := question.NewGetQuestionByIdHandler(&usecase.QuestionService)

	router.Group(func(r chi.Router) {
		r.Use(SetupHeaders)
		r.Get("/question_random/", questionRandomHandler.ServeHTTP)
		r.Get("/question/{id}/", questionByIdHandler.ServeHTTP)
	})
}

func SetupHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
