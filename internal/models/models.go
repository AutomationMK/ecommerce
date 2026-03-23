package models

import (
	"time"

	"github.com/jackc/pgx/v5"
)

// DBModel is the type for database connection values
type DBModel struct {
	DB *pgx.Conn
}

// Models is the wrapper for all models
type Models struct {
	DB DBModel
}

// NewModels returns a model type with pgx connection pool
func NewModels(db *pgx.Conn) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Widget is the type for all widgets
type Widget struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	InventoryLevel int       `json:"inventory_level"`
	Price          int       `json:"price"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}
