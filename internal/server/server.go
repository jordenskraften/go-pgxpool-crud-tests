package server

import (
	"log"
	"log/slog"
	handler "pxgpool-crud-tests/internal/server/handler"
	"time"

	"net/http"
)

type Server struct {
	Server *http.Server
	Router *http.ServeMux
	logger *slog.Logger
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
		logger: logger,
	}
}
func (srv *Server) Start() {
	go log.Fatal(srv.Server.ListenAndServe())
}
func (srv *Server) RegisterHandler(handler handler.HttpHandler) {
	srv.Router.HandleFunc(handler.GetUrlPattern(), handler.GetHandler())
}

func InitRouter() *http.ServeMux {

	router := http.NewServeMux()

	return router
}
