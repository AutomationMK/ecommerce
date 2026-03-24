package models

import (
	"context"
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
	Image          string    `json:"image"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}

// GetWidget gets widget data from the database
func (m *DBModel) GetWidget(id int) (Widget, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var widget Widget

	query := `
		SELECT id, name, description, inventory_level, price, created_at, updated_at
		FROM widgets
		WHERE id = $1
	`
	row := m.DB.QueryRow(ctx, query, id)
	err := row.Scan(&widget.ID, &widget.Name, &widget.Description,
		&widget.InventoryLevel, &widget.Price, &widget.CreatedAt, &widget.UpdatedAt)
	if err != nil {
		return widget, err
	}

	return widget, nil
}

// Order is the type for all orders
type Order struct {
	ID            int       `json:"id"`
	WidgetId      int       `json:"widget_id"`
	TransactionId int       `json:"transaction_id"`
	StatusId      int       `json:"status_id"`
	Quantity      int       `json:"quantity"`
	Amount        int       `json:"amount"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

// Status is the type for all statuses
type Status struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// TransactionStatus is the type for all transaction statuses
type TransactionStatus struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Transaction is the type for all transactions
type Transaction struct {
	ID                  int       `json:"id"`
	Amount              int       `json:"amount"`
	Currency            string    `json:"currency"`
	LastFour            string    `json:"last_four"`
	BankReturnCode      string    `json:"bank_return_code"`
	TransactionStatusId int       `json:"transaction_status_id"`
	CreatedAt           time.Time `json:"-"`
	UpdatedAt           time.Time `json:"-"`
}

// User is the type for all users
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
