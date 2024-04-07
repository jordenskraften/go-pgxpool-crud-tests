package repository

import (
	database "pxgpool-crud-tests/internal/db"
	authRepo "pxgpool-crud-tests/internal/repository/auth"
	questRepo "pxgpool-crud-tests/internal/repository/question"
	"pxgpool-crud-tests/internal/usecase/auth"
	"pxgpool-crud-tests/internal/usecase/question"
)

// здесь будут инициализироваться все репо, чтобы удобно пробросить их дальше
type Repository struct {
	AuthRepo     auth.AuthRepository
	QuestionRepo question.QuestionRepository
}

func NewRepository(db *database.DB) *Repository {
	r := &Repository{
		AuthRepo:     authRepo.NewAuthRepositoryMock(),
		QuestionRepo: questRepo.NewQuestionRepositoryPostgres(db),
	}
	return r
}
