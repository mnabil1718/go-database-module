package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// connection string format: postgres://username:password@localhost:5432/database_name
	db, err := sql.Open("pgx", "postgres://postgres:Cucibaju123@localhost:5432/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var greeting string
	err = db.QueryRow("select 'hello world'").Scan(&greeting)
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)
}
