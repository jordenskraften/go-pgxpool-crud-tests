package server

import (
	"log/slog"
	"time"

	"net/http"
)

type Server struct {
	Server *http.Server
	Router *http.ServeMux
}

func NewServer(logger *slog.Logger) *Server {
	router := InitRouter()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		Server: server,
		Router: router,
	}
}

func InitRouter() *http.ServeMux {

	router := http.NewServeMux()

	return router
}
