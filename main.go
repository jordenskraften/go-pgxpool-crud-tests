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
	"pxgpool-crud-tests/internal/service"
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
	repo := repository.NewRepository(logger, db)
	exampleService := service.NewExampleService(logger, repo)

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

/*
func insertsTests(){

	// Database queries
	logger.Info("starting inserting performing tests")
	const totalInserts int = 100
	const threadsNumber int = 10
	jobs := make(chan int, totalInserts)

	// Предположим, что connPool уже инициализирован

	var wg sync.WaitGroup

	// Инициализация воркеров
	for w := 1; w <= threadsNumber; w++ {
		wg.Add(1)
		go worker(ctx, db.Pool, logger, jobs, &wg, w)
	}

	// Отправка работы воркерам
	go func() {
		for j := 1; j <= totalInserts; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	go func() {
		time.Sleep(time.Second * 5)
		cancel()
	}()
	// Ожидание завершения всех воркеров
	wg.Wait()
	logger.Info("inserts are done")

	//-------------------------------
}
func worker(ctx context.Context, connPool *pgxpool.Pool, logger *slog.Logger, jobs <-chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()

loop:
	for {
		select {
		case val := <-jobs:
			if val <= 0 {
				break loop
			}
			// Воркер читает счетчик в канале до тех пор, пока он не опустеет до нуля
			logger.Info(fmt.Sprintf("inside in worker %d counter = %d", id, val))
			err := InsertQuery(connPool, logger)
			if err != nil {
				logger.Error(err.Error())
				break loop
			}
		case <-ctx.Done():
			logger.Info(fmt.Sprintf("context is canceled id worker is %d", id))
			break loop
		}
	}
}

func InsertQuery(p *pgxpool.Pool, logger *slog.Logger) error {
	//эт репозиторий  верней конкретный его метод
	// handler -> concrete N handlers
	// repo -> concrete N repo
	title := RandomCryptoString(10, logger)
	question := RandomCryptoString(20, logger)
	//trying to SQL INJECT
	anwser := "Delete * from Question" + RandomCryptoString(80, logger)
	_, err := p.Exec(context.Background(),
		"insert into Question(question_title, question_text, answer_text) values($1, $2, $3)",
		title, question, anwser)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

func RandomCryptoString(length int, logger *slog.Logger) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		logger.Error(err.Error())
		return "" // В случае ошибки возвращаем её
	}

	// Кодируем полученные байты в строку с помощью base64
	return base64.URLEncoding.EncodeToString(b)
}
*/
