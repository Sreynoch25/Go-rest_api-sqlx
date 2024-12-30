package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)


/*
 * Author:Noch
 * ConnectDB initializes a connection to the PostgreSQL database using the `DATABASE_URL` environment variable.
 */
func ConnectDB() (*sqlx.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is required but not set")
	}

	// Connect to the PostgreSQL database using the connection URL
	db, err := sqlx.Connect("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	fmt.Println("Database connected!")
	return db, nil
}
