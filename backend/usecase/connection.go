package usecase

import (
	"capydb/backend/dto/request"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (u *Usecase) CreateConnection(req request.CreateConnection) error {
	// Only supports PostgreSQL for now
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", req.User, req.Password, req.Host, req.Port, req.DatabaseName)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return err
	}
	u.dbConn = dbConn

	return nil
}
