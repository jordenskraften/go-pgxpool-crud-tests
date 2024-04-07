package server

import (
	"log"
	"log/slog"
	"pxgpool-crud-tests/internal/usecase"
	"time"

	"net/http"
)

type Server struct {
	Server *http.Server
	logger *slog.Logger
}

func NewServer(usecase *usecase.Usecase, logger *slog.Logger) *Server {
	router := NewRouter(usecase)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		Server: server,
		logger: logger,
	}
}
func (srv *Server) Start() {
	go log.Fatal(srv.Server.ListenAndServe())
}
