package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	usecaseAuth "pxgpool-crud-tests/internal/usecase/auth"
	"time"
)

//todo типичный роут который дергает бизнес логику авторизации

type AuthHandler struct {
	PatternUrl  string
	usecase     *usecaseAuth.AuthService
	HttpHandler func(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler(patternUrl string, usecase *usecaseAuth.AuthService) *AuthHandler {
	handler := &AuthHandler{
		PatternUrl:  patternUrl,
		usecase:     usecase,
		HttpHandler: nil,
	}
	handler.HttpHandler = handler.ServeHTTP

	return handler
}

func (ah *AuthHandler) GetUrlPattern() string {
	return ah.PatternUrl
}
func (ah *AuthHandler) GetHandler() func(http.ResponseWriter, *http.Request) {
	return ah.HttpHandler
}

func (ah *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")
	password := r.URL.Query().Get("password")
	log.Println(login)
	log.Println(password)
	token := ""
	if login != "" && password != "" {
		token = fmt.Sprintf("%s:%s", login, password)
	}

	log.Println(token)
	ok, err := ah.usecase.AuthUser(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "User validation error", http.StatusUnauthorized)
		return
	}

	// Сохранение токена в куки
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:    "auth_token",
		Value:   token,
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)

	// Подготовка и отправка ответа клиенту
	type AuthResponse struct {
		Message string `json:"auth_message"`
	}
	response := &AuthResponse{
		Message: "auth is success!",
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
