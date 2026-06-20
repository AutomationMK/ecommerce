package models

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// DBModel is the type for database connection values
type DBModel struct {
	DB *pgxpool.Pool
}

// Models is the wrapper for all models
type Models struct {
	DB DBModel
}

// NewModels returns a model type with pgx connection pool
func NewModels(db *pgxpool.Pool) Models {
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
	IsRecurring    bool      `json:"is_recurring"`
	PlanID         string    `json:"plan_id"`
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
			is_recurring, plan_id, created_at, updated_at
		FROM widgets
		WHERE id = $1
	`
	row := m.DB.QueryRow(ctx, query, id)
	err := row.Scan(&widget.ID, &widget.Name, &widget.Description,
		&widget.InventoryLevel, &widget.Price, &widget.Image,
		&widget.IsRecurring, &widget.PlanID, &widget.CreatedAt,
		&widget.UpdatedAt)
	if err != nil {
		return widget, err
	}

	return widget, nil
}

// GetWidget gets widget data from the database
func (m *DBModel) GetWidgetByName(name string) (Widget, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var widget Widget

	query := `
		SELECT id, name, description, inventory_level, price, image,
			is_recurring, plan_id, created_at, updated_at
		FROM widgets
		WHERE name = $1
	`
	row := m.DB.QueryRow(ctx, query, name)
	err := row.Scan(&widget.ID, &widget.Name, &widget.Description,
		&widget.InventoryLevel, &widget.Price, &widget.Image,
		&widget.IsRecurring, &widget.PlanID, &widget.CreatedAt,
		&widget.UpdatedAt)
	if err != nil {
		return widget, err
	}

	return widget, nil
}

// Order is the type for all orders
type Order struct {
	ID            int         `json:"id"`
	WidgetID      int         `json:"widget_id"`
	TransactionID int         `json:"transaction_id"`
	StatusID      int         `json:"status_id"`
	CustomerID    int         `json:"customer_id"`
	Quantity      int         `json:"quantity"`
	Amount        int         `json:"amount"`
	Widget        Widget      `json:"widget"`
	Transaction   Transaction `json:"transaction"`
	Customer      Customer    `json:"customer"`
	CreatedAt     time.Time   `json:"-"`
	UpdatedAt     time.Time   `json:"-"`
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

// GetAllOrders returns all orders that are non-recurring
func (m *DBModel) GetAllOrders() ([]*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var orders []*Order
	query := `
		SELECT o.id, o.status_id, o.quantity, o.created_at, o.updated_at,
			o.widget_id, w.id, w.name, o.transaction_id, t.id, t.amount,
			t.currency, t.last_four, t.expiry_month, t.expiry_year,
			t.payment_intent, t.payment_method, t.bank_return_code, o.customer_id, c.id,
			c.first_name, c.last_name, c.email
		FROM orders AS o
			LEFT JOIN widgets AS w ON o.widget_id = w.id
			LEFT JOIN transactions AS t ON o.transaction_id = t.id
			LEFT JOIN customers AS c ON o.customer_id = c.id
		WHERE w.is_recurring = false
		ORDER BY o.created_at DESC`
	rows, err := m.DB.Query(ctx, query)
	if err != nil {
		return orders, err
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		err = rows.Scan(
			&order.ID,
			&order.StatusID,
			&order.Quantity,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.WidgetID,
			&order.Widget.ID,
			&order.Widget.Name,
			&order.TransactionID,
			&order.Transaction.ID,
			&order.Transaction.Amount,
			&order.Transaction.Currency,
			&order.Transaction.LastFour,
			&order.Transaction.ExpiryMonth,
			&order.Transaction.ExpiryYear,
			&order.Transaction.PaymentIntent,
			&order.Transaction.PaymentMethod,
			&order.Transaction.BankReturnCode,
			&order.CustomerID,
			&order.Customer.ID,
			&order.Customer.FirstName,
			&order.Customer.LastName,
			&order.Customer.Email,
		)
		if err != nil {
			return orders, err
		}

		orders = append(orders, &order)
	}
	if err = rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil
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

// GetUserByEmail gets a user by email address
func (m *DBModel) GetUserByEmail(email string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	email = strings.ToLower(email)
	var u User

	row := m.DB.QueryRow(ctx, `
		SELECT id, first_name, last_name, email, password,
			created_at, updated_at
		FROM users
		WHERE email = $1`, email)
	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}

// Authenticate compares the hashes of input password and database password
// if success then return the user id else return error
func (m *DBModel) Authenticate(email, password string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	var hashedPassword string

	row := m.DB.QueryRow(ctx, `
		SELECT id, password
		FROM users
		WHERE email = $1`, email)
	if err := row.Scan(&id, &hashedPassword); err != nil {
		return id, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, errors.New("incorrect password")
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdatePasswordForUser updates a users password hash in the database
func (m *DBModel) UpdatePasswordForUser(u User, hash string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		UPDATE users
		SET password = $1
		WHERE id = $2`
	_, err := m.DB.Exec(ctx, stmt, hash, u.ID)
	if err != nil {
		return err
	}

	return nil
}
