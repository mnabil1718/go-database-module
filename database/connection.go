package database

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetConnection() *sql.DB {
	// connection string format: postgres://username:password@localhost:5432/database_name
	db, err := sql.Open("pgx", "postgres://postgres:Cucibaju123@localhost:5432/go_db_module")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db

}
