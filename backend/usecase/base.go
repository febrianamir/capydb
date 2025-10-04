package usecase

import "gorm.io/gorm"

type Usecase struct {
	dbConn *gorm.DB
}

func NewUsecase() *Usecase {
	return &Usecase{}
}
