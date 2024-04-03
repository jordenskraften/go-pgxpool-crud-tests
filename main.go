package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	database "pxgpool-crud-tests/internal/db"
	"pxgpool-crud-tests/internal/logger"
	repoQuestion "pxgpool-crud-tests/internal/repository/question"
	"pxgpool-crud-tests/internal/server"
	"pxgpool-crud-tests/internal/server/handlers"
	usecaseQuestion "pxgpool-crud-tests/internal/usecase/question"
	"syscall"
)

func main() {
	logger := logger.NewLogger()
	logger.Info("app is started")
	logger.Info("httpmuxgo121 must be = 0 for new 1.22 http mux:", slog.String("GODEBUG ", os.Getenv("GODEBUG")))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := database.NewDb(ctx, logger)
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}
	defer db.Pool.Close()

	httpServer := server.NewServer(logger)

	questionRepo := repoQuestion.NewQuestionRepositoryPostgres(db)
	questionUsecase := usecaseQuestion.NewQuestionService(questionRepo)
	questionHandler := handlers.NewGetQuestionHandler(httpServer.Router, questionUsecase, logger)

	httpServer.Router.HandleFunc("/question", questionHandler.ServeHTTP)
	httpServer.Router.HandleFunc("GET /quest/", questionHandler.ServeHTTP)
	httpServer.Router.HandleFunc("GET /path/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("got path\n")
	})

	go httpServer.Server.ListenAndServe()
	// Инициализация грациозного завершения работы
	gracefulShutdown(logger)
}

func gracefulShutdown(logger *slog.Logger) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	logger.Warn("Graceful Shutdown Initiated")
	os.Exit(0)
}
