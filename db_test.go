package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

func TestQueryColumnType(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, `INSERT INTO
	customer(id, name, email, balance, rating, birth_date, married)
	VALUES ('CUST01', 'Joko Widodo', 'joko@gmail.com', 6000000, 4.9, '1974-4-12', true),
	('CUST02', 'Susilo Bambang', 'susilo@gmail.com', 5000000, 4.8, '1975-5-13', false),
	('CUST03', 'Megawati Soekarno', 'megawati@gmail.com', 7000000, 4.7, '1976-6-14', true);`)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Insert data success")

	rows, err := db.QueryContext(ctx, "SELECT id, name, email, password, balance, rating, created_at, birth_date, married FROM customer;")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id, name, password string
		var email sql.NullString
		var balance sql.NullInt32
		var rating sql.NullFloat64
		var birthDate, created_at sql.NullTime
		var married sql.NullBool
		err := rows.Scan(&id, &name, &email, &password, &balance, &rating, &created_at, &birthDate, &married)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		} else {
			fmt.Println("Email: NULL")
		}
		fmt.Println("Password:", password)
		if balance.Valid {
			fmt.Println("Balance:", balance.Int32)
		} else {
			fmt.Println("Balance: NULL")
		}
		if rating.Valid {
			fmt.Println("Rating:", rating.Float64)
		} else {
			fmt.Println("Rating: NULL")
		}
		if created_at.Valid {
			fmt.Println("Created At:", created_at.Time)
		} else {
			fmt.Println("Created At: NULL")
		}
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		} else {
			fmt.Println("Birth Date: NULL")
		}
		if married.Valid {
			fmt.Println("Married:", married.Bool)
		} else {
			fmt.Println("Married: NULL")
		}
		fmt.Println("====================================================================")
	}
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	statement, err := db.PrepareContext(ctx, "INSERT INTO comments (email, comment) VALUES ($1, $2) RETURNING id;")
	if err != nil {
		panic(err.Error())
	}
	defer statement.Close()

	for i := range 10 {
		email := "joko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke-" + strconv.Itoa(i)

		var lastIndex int
		err = statement.QueryRowContext(ctx, email, comment).Scan(&lastIndex)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Insert data success with id:", lastIndex)
	}
}
