package db

import (
	"database/sql"
)

type Store interface {
	Querier
}

// SQLStore provides all functions to execute db queries and transactions
type SQLStore struct {
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
	}
}
