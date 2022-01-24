package database

import "wazomera/app/queries"

// Queries struct - collects all app queries.
type Queries struct {
	*queries.BookQueries // load queries from the Book model
}

// OpenDBConnection func - opens database connection.
func OpenDBConnection() (*Queries, error) {
	// a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		BookQueries: &queries.BookQueries{DB: db}, // from Book model
	}, nil
}