package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log/slog"
	"pxgpool-crud-tests/internal/logger"
	"strconv"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	logger := logger.NewLogger()
	logger.Info("app is started")

	// Create database connection
	connPool, err := pgxpool.NewWithConfig(context.Background(), ConfigDB(logger))
	if err != nil {
		logger.Error(err.Error())
	}

	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		logger.Error(err.Error())
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("connected to database")

	// Database queries
	wg := &sync.WaitGroup{}
	counter := 100000
	//1 mln parallel insert queries

	logger.Info("starting inserting performing tests")
	for i := 0; i < counter; i++ {
		wg.Add(1)
		go InsertQuery(connPool, logger, wg)
		/*
			SELECT
				(SELECT COUNT(*) FROM Question) AS total_records,
				answer_text
			FROM
			Question
			LIMIT 10;
		*/
	}
	wg.Wait()
	time.Sleep(2 * time.Second)
	logger.Info("inserting tests is over")

	time.Sleep(10 * time.Second)
	defer connPool.Close()
	//это будет долго
	//крч пул подключений дергается и паралельно записи идут
	//даже на такой нагрузке не падает
	//ладно миллион не буду ждать, останавливаю
}
func InsertQuery(p *pgxpool.Pool, logger *slog.Logger, wg *sync.WaitGroup) {
	defer wg.Done()

	title := RandomCryptoString(10, logger)
	question := RandomCryptoString(20, logger)
	//trying to SQL INJECT
	anwser := "Delete * from Question where question_title != " + RandomCryptoString(80, logger)
	_, err := p.Exec(context.Background(),
		"insert into Question(question_title, question_text, answer_text) values($1, $2, $3)",
		title, question, anwser)
	if err != nil {
		logger.Error(err.Error())
	}
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

func ConfigDB(logger *slog.Logger) *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	// Your own Database URL
	dbUser := "postgres"
	dbPassword := "mysecretpassword"
	dbHost := "localhost"
	dbPort := strconv.Itoa(5432) // Преобразование порта из int в string
	dbName := "interview_spellbook"
	// Исправление с добавлением плейсхолдеров для значений
	dbAddress := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?", dbUser, dbPassword, dbHost, dbPort, dbName)

	dbConfig, err := pgxpool.ParseConfig(dbAddress)
	if err != nil {
		logger.Error(err.Error())
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		logger.Info("Before acquiring the connection pool to the database!!")

		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		logger.Info("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		logger.Info("Closed the connection pool to the database!!")
	}

	return dbConfig
}
