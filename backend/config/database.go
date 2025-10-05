package config

import (
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(appDataDir string) (*gorm.DB, error) {
	// Ensure app data directory exists
	if err := os.MkdirAll(appDataDir, 0755); err != nil {
		return nil, err
	}

	// Place sqlite file inside app's data dir
	dbPath := filepath.Join(appDataDir, "capydb.db")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Create credential database
	schema := `
	CREATE TABLE IF NOT EXISTS credentials (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		db_vendor TEXT NOT NULL,
		hex_color TEXT NOT NULL,
		host TEXT NOT NULL,
		port TEXT NOT NULL,
		user TEXT NOT NULL,
		password TEXT NOT NULL,
		database_name TEXT NOT NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL
	);
	`
	if result := db.Exec(schema); result.Error != nil {
		return nil, result.Error
	}

	return db, nil
}
