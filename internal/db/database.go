package database

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbConnectAttemptions = 10
	dbConnectAttemptTime = time.Second
)

type DB struct {
	logger *slog.Logger
	Pool   *pgxpool.Pool
}

func NewDb(ctx context.Context, logger *slog.Logger) (*DB, error) {
	for i := 0; i < dbConnectAttemptions; i++ {
		connPool, err := pgxpool.NewWithConfig(ctx, configDB(logger))
		if err != nil {
			logger.Error("Failed to create connection pool: ", err)
			time.Sleep(dbConnectAttemptTime)
			continue
		}

		// Attempt to acquire a connection
		conn, err := connPool.Acquire(ctx)
		if err != nil {
			logger.Error("Failed to acquire connection: ", err)
			time.Sleep(dbConnectAttemptTime)
			continue
		}
		conn.Release() // Release the connection as soon as we're done checking

		// Successfully connected
		db := &DB{
			logger: logger,
			Pool:   connPool,
		}
		logger.Info("Successfully connected to database.")
		return db, nil
	}

	logger.Error("Unable to connect to database after several attempts.")
	return nil, fmt.Errorf("unable to connect to database after %d attempts", dbConnectAttemptions)
}

func configDB(logger *slog.Logger) *pgxpool.Config {
	const defaultMaxConns = int32(5)
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
		//logger.Info("Before acquiring the connection pool to the database!!")

		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		//logger.Info("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		logger.Info(fmt.Sprintf("Closed the connection pool to the database with id %d", c))
	}

	return dbConfig
}
