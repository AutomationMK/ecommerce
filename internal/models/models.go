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
		SELECT id, name, description, inventory_level, price, image,
			created_at, updated_at
		FROM widgets
		WHERE id = $1
	`
	row := m.DB.QueryRow(ctx, query, id)
	err := row.Scan(&widget.ID, &widget.Name, &widget.Description,
		&widget.InventoryLevel, &widget.Price, &widget.Image,
		&widget.CreatedAt, &widget.UpdatedAt)
	if err != nil {
		return widget, err
	}

	return widget, nil
}

// Order is the type for all orders
type Order struct {
	ID            int       `json:"id"`
	WidgetID      int       `json:"widget_id"`
	TransactionID int       `json:"transaction_id"`
	StatusID      int       `json:"status_id"`
	CustomerID    int       `json:"customer_id"`
	Quantity      int       `json:"quantity"`
	Amount        int       `json:"amount"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

// InsertOrder inserts a new order into the database
func (m *DBModel) InsertOrder(ord Order) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO orders (
			widget_id, transaction_id, status_id, quantity,
			amount, customer_id, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	var newID int
	err := m.DB.QueryRow(ctx, stmt,
		ord.WidgetID,
		ord.TransactionID,
		ord.StatusID,
		ord.Quantity,
		ord.Amount,
		ord.CustomerID,
		time.Now().UTC(),
		time.Now().UTC(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
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
	TransactionStatusID int       `json:"transaction_status_id"`
	ExpiryMonth         int       `json:"expiry_month"`
	ExpiryYear          int       `json:"expiry_year"`
	PaymentIntent       string    `json:"payment_intent"`
	PaymentMethod       string    `json:"payment_method"`
	CreatedAt           time.Time `json:"-"`
	UpdatedAt           time.Time `json:"-"`
}

// InsertTransaction inserts a new transaction into the database
func (m *DBModel) InsertTransaction(txn Transaction) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO transactions (
			amount, currency, last_four, bank_return_code, expiry_month,
			expiry_year, transaction_status_id, payment_intent, payment_method,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`

	var newID int
	err := m.DB.QueryRow(ctx, stmt,
		txn.Amount,
		txn.Currency,
		txn.LastFour,
		txn.BankReturnCode,
		txn.ExpiryMonth,
		txn.ExpiryYear,
		txn.TransactionStatusID,
		txn.PaymentIntent,
		txn.PaymentMethod,
		time.Now().UTC(),
		time.Now().UTC(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
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

// Customer is the type for all customers
type Customer struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// InsertCustomer inserts a new customer into the database
func (m *DBModel) InsertCustomer(cus Customer) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO customers (
			first_name, last_name, email,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var newID int
	err := m.DB.QueryRow(ctx, stmt,
		cus.FirstName,
		cus.LastName,
		cus.Email,
		time.Now().UTC(),
		time.Now().UTC(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}
