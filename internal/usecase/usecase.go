package usecase

import (
	"pxgpool-crud-tests/internal/repository"
	"pxgpool-crud-tests/internal/usecase/auth"
	"pxgpool-crud-tests/internal/usecase/question"
)

type Usecase struct {
	AuthService     auth.AuthService
	QuestionService question.QuestionService
}

func NewUsecase(repo *repository.Repository) *Usecase {
	u := &Usecase{
		AuthService:     *auth.NewAuthService(repo.AuthRepo),
		QuestionService: *question.NewQuestionService(repo.QuestionRepo),
	}
	return u
}
