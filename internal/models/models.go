package models

import "github.com/jackc/pgx/v5"

// DBModel is the type for database connection values
type DBModel struct {
	DB *pgx.Conn
}
