package config

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func InitDB(appDataDir string) (*sql.DB, error) {
	// Ensure app data directory exists
	if err := os.MkdirAll(appDataDir, 0755); err != nil {
		return nil, err
	}

	// Place sqlite file inside app's data dir
	dbPath := filepath.Join(appDataDir, "capydb.db")

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	// Create credential database
	schema := `
	CREATE TABLE IF NOT EXISTS credentials (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		host TEXT NOT NULL,
		port TEXT NOT NULL,
		user TEXT NOT NULL,
		password TEXT NOT NULL,
		database_name TEXT NOT NULL
	);
	`
	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}
