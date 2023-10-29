package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type dao struct {}

type DAO interface {
	NewAppRepository() AppHealthRepository
}

var DB *pgxpool.Pool


func InitConn() (DAO) {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	// defer DB.Close()

	var greeting string
	err = pool.QueryRow(context.Background(), "select 'Connected!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
	DB = pool
	return &dao{}
}

func (d *dao) NewAppRepository() AppHealthRepository {
	return &appHealthRepository{}
}