package main

import (
	"context"
	"database/sql"
	"fmt"
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

func main() {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name) VALUES('joko', 'Joko');")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Insert data success")

	rows, err := db.Query("SELECT * FROM customer;")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("ID:", id, "Name:", name)
	}

}
