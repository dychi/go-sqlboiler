package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	dsn := fmt.Sprintf("dbname=%s user=%s",
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_USER"),
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return db
}
