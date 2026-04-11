package service

import "gorm.io/gorm"

type UserDataRepository interface {
	CreateConfig(cfg *Model) error
	GetUserById(id int) (*Model, error)
	GetUsersList() ([]Model, error)
	GetOrderById(orderId int) (*Model, error)
	GetOrdersList() ([]Model, error)
}

type DBState struct {
	db *gorm.DB
}

func NewCfgRepository(db *gorm.DB) UserDataRepository {
	return &DBState{db: db}
}

func (s *DBState) CreateConfig(cfg Model) error {
	return s.db.Create(&cfg).Error
}
