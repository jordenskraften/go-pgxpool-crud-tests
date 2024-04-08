package server

import (
	"pxgpool-crud-tests/internal/server/handler/question"
	"pxgpool-crud-tests/internal/server/middleware"
	"pxgpool-crud-tests/internal/usecase"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewRouter(usecase *usecase.Usecase) chi.Router {
	r := chi.NewRouter()
	InitRouter(usecase, r)
	return r
}

func InitRouter(usecase *usecase.Usecase, r chi.Router) {
	questionRandomHandler := question.NewGetRandomQuestionHandler(&usecase.QuestionService)
	questionByIdHandler := question.NewGetQuestionByIdHandler(&usecase.QuestionService)
	addQustionHandler := question.NewAddQuestionHandler(&usecase.QuestionService)

	//публичные

	r.Group(func(r chi.Router) {
		r.Use(middleware.SetupHeaders)
		r.Use(chiMiddleware.StripSlashes) // Добавить StripSlashes здесь
		r.Get("/question_random", questionRandomHandler.ServeHTTP)
		r.Get("/question_random/", questionRandomHandler.ServeHTTP)
		r.Get("/question/{id}", questionByIdHandler.ServeHTTP)
		r.Get("/question/{id}/", questionByIdHandler.ServeHTTP)
	})

	//приватные
	r.Group(func(r chi.Router) {
		r.Use(middleware.SetupHeaders)
		r.Use(middleware.AuthMiddleware)
		r.Post("/add_question", addQustionHandler.ServeHTTP)
		r.Post("/add_question/", addQustionHandler.ServeHTTP)
	})
}
