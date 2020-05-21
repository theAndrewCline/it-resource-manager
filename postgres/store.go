package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	// tutorial told me to do it this way
	_ "github.com/lib/pq"
)

// NewStore creates a database connection and returns a store
func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("Error opening the database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Error connecting to the database: %w", err)
	}

	return &Store{
		OwnerStore: &OwnerStore{
			DB: db,
		},
		ComputerStore: &ComputerStore{
			DB: db,
		},
		PartStore: &PartStore{
			DB: db,
		},
	}, nil
}

// Store implements the struct for the store
type Store struct {
	*OwnerStore
	*ComputerStore
	*PartStore
}
