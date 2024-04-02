package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	database "pxgpool-crud-tests/internal/db"
	"pxgpool-crud-tests/internal/logger"
	"pxgpool-crud-tests/internal/repository"
	"pxgpool-crud-tests/internal/usecase"
	"syscall"
)

func main() {
	logger := logger.NewLogger()
	logger.Info("app is started")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := database.NewDb(ctx, logger)
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}
	defer db.Pool.Close()
	repo := repository.NewRepositoryPostgres(logger, db)
	exampleService := usecase.NewExampleService(logger, repo)

	logger.Info("starting 10 get 10 random questions")
	for i := 0; i < 10; i++ {
		quest, err := exampleService.GetRandomQuestion()
		if err != nil {
			logger.Error("failed while execution get random question service", err)
		}
		logger.Info("get random question query is good, value:", quest.Answer_text)
	}
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
