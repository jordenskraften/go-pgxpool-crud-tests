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
	"pxgpool-crud-tests/internal/repository"
	"pxgpool-crud-tests/internal/server"
	"pxgpool-crud-tests/internal/usecase"
	"syscall"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil || cookie.Value == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Продолжаем выполнение следующего обработчика
		next.ServeHTTP(w, r)
	})
}

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

	repoCommon := repository.NewRepository(db)
	usecaseCommon := usecase.NewUsecase(repoCommon)
	httpServer := server.NewServer(usecaseCommon, logger)
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
