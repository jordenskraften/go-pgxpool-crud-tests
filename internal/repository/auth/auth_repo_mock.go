package auth

import "errors"

type AuthRepositoryMock struct {
}

func NewAuthRepositoryMock() *AuthRepositoryMock {
	return &AuthRepositoryMock{}
}

func (a AuthRepositoryMock) ValidateUser(token string) (bool, error) {
	if token != "" {
		return true, nil
	}
	return false, errors.New("user auth validation fail in repo stage!")
}
