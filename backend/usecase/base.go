package usecase

import (
	"capydb/backend/config"
	"context"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

type Usecase struct {
	activeDbConns map[string]*gorm.DB
	dbDataConn    *gorm.DB
}

func (u *Usecase) getActiveDbConn(connectionId string) *gorm.DB {
	if dbConn, ok := u.activeDbConns[connectionId]; ok {
		return dbConn
	}
	return nil
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
	usecase.activeDbConns = map[string]*gorm.DB{} // Ensure the connections map is not nil
	return usecase
}
