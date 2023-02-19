package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

const AppName = "Events"

func newDBConnection() (*pgxpool.Pool, error) {
	min := 3
	max := 100
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	urlDB := makeDNS(user, pass, host, port, dbName, sslMode, min, max)

	config, err := pgxpool.ParseConfig(urlDB)

	if err != nil {
		return nil, fmt.Errorf("%s %w", "pgxpool.ParseConfig", err)
	}

	config.ConnConfig.RuntimeParams["application_name"] = AppName

	pool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		return nil, fmt.Errorf("%s %w", "pgxpool.NewWithConfig", err)
	}

	return pool, nil

}

func makeDNS(user, pass, host, port, dbName, sslMode string, minConn, maxConn int) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_min_conns=%d pool_max_conns=%d",
		user,
		pass,
		host,
		port,
		dbName,
		sslMode,
		minConn,
		maxConn)
}
