package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	dsn := "host=localhost port=5432 user=user password=password dbname=aidvault sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to PostgreSQL")
	return db, nil
}
