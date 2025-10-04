package usecase

import (
	"capydb/backend/config"
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

type Usecase struct {
	dbConn     *gorm.DB
	dbDataConn *sql.DB
}

func appDataDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(base, "capydb")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return dir, nil
}

func NewUsecase(ctx context.Context) *Usecase {
	usecase := &Usecase{}

	appDataDir, err := appDataDir()
	if err != nil {
		runtime.LogError(ctx, err.Error())
		return nil
	}

	db, err := config.InitDB(appDataDir)
	if err != nil {
		runtime.LogError(ctx, err.Error())
		return nil
	}

	usecase.dbDataConn = db
	return &Usecase{}
}
