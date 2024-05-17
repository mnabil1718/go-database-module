package entity

import "database/sql"

type Comment struct {
	Id      int64
	Email   string
	Comment sql.NullString
}
