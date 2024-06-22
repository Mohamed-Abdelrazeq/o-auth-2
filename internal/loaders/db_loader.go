package loaders

import (
	"Mohamed-Abdelrazeq/o-auth-2/internal/database"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseInstance struct {
	DB *database.Queries
}

func LoadDB() *DatabaseInstance {
	db, err := sql.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}

	queries := database.New(db)

	// Init DB Instance
	return &DatabaseInstance{queries}
}
