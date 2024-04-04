package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	database "pxgpool-crud-tests/internal/db"
	"pxgpool-crud-tests/internal/logger"
	authRepo "pxgpool-crud-tests/internal/repository/auth"
	questionRepo "pxgpool-crud-tests/internal/repository/question"
	"pxgpool-crud-tests/internal/server"
	authHandler "pxgpool-crud-tests/internal/server/handler/auth"
	questionHandler "pxgpool-crud-tests/internal/server/handler/question"
	authUsecase "pxgpool-crud-tests/internal/usecase/auth"
	questionUsecase "pxgpool-crud-tests/internal/usecase/question"
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

	questionRepo := questionRepo.NewQuestionRepositoryPostgres(db)
	questionUsecase := questionUsecase.NewQuestionService(questionRepo)
	randomQuestionHandler := questionHandler.NewGetRandomQuestionHandler("GET /question_random/", questionUsecase)
	questionByIdHandler := questionHandler.NewGetQuestionByIdHandler("GET /question/{id}/", questionUsecase)
	httpServer.RegisterHandler(randomQuestionHandler)
	httpServer.RegisterHandler(questionByIdHandler)

	authRepo := authRepo.NewAuthRepositoryMock()
	authUsecase := authUsecase.NewAuthService(authRepo)
	authHandler := authHandler.NewAuthHandler("POST /auth/", authUsecase)
	httpServer.RegisterHandler(authHandler)
	//как-то надо все роуты регать нормально, а не так вот...

	//листен надо бы внутрь структуры сервер вставить с остановкой по контексту и сигналу хз
	httpServer.Start()
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
