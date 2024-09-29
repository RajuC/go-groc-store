package database

import (
	"database/sql"
)

// Service represents a service that interacts with a database.
type (
	Service interface {
		// Health returns a map of health status information.
		// The keys and values in the map are service-specific.
		Health() map[string]string
		// New(l *slog.Logger, cfg *config.Config) *DbClient

		// Close terminates the database connection.
		// It returns an error if the connection cannot be closed.
		Close() error

		CreateCustomerTable()
		CreateCustomer(customer []byte) (sql.Result, error)
		GetCustomer(id int64) *sql.Row
		ListAllCustomers() (*sql.Rows, error)
	}
)
