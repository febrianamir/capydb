package usecase

import (
	"capydb/backend/dto/request"
	"capydb/backend/dto/response"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (u *Usecase) CreateConnection(req request.CreateConnection) (response.CreateConnection, error) {
	// Create connection id
	connectionId := uuid.NewString()

	// Only supports PostgreSQL for now
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", req.User, req.Password, req.Host, req.Port, req.DatabaseName)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return response.CreateConnection{}, err
	}

	// Add new connection to activeDbConns
	u.activeDbConns[connectionId] = dbConn
	return response.CreateConnection{ConnectionId: connectionId}, nil
}
