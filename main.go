package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	log.Println("app is started")

	// Create database connection
	connPool, err := pgxpool.NewWithConfig(context.Background(), Config())
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}

	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from the database pool!!")
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatal("Could not ping database")
	}

	fmt.Println("Connected to the database!!")

	// Database queries
	InsertQuery(connPool)
	time.Sleep(time.Second)
	InsertQuery(connPool)
	InsertQuery(connPool)

	defer connPool.Close()

}
func InsertQuery(p *pgxpool.Pool) {
	_, err := p.Exec(context.Background(),
		"insert into Question(question_title, question_text, answer_text) values($1, $2, $3)",
		RandomCryptoString(10), RandomCryptoString(20), RandomCryptoString(80))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func RandomCryptoString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "" // В случае ошибки возвращаем её
	}

	// Кодируем полученные байты в строку с помощью base64
	return base64.URLEncoding.EncodeToString(b)
}

func Config() *pgxpool.Config {
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
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	return dbConfig
}
