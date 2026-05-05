package repository

import (
	mod "rs/internal/models"
	serv "rs/internal/service"

	"gorm.io/gorm"
)

type DBStateUsers struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) serv.UserDataRepository {
	return &DBStateUsers{db: db}
}

func (r *DBStateUsers) CreateUser(user *mod.CfgVpnModel) error {
	return r.db.Create(&user).Error
}

func (r *DBStateUsers) GetUserById(id int) (mod.CfgVpnModel, error) {
	var user mod.CfgVpnModel

	err := r.db.First(&user, id).Error

	return user, err
}

func (r *DBStateUsers) GetUsersList() ([]mod.CfgVpnModel, error) {
	var user []mod.CfgVpnModel

	return user, r.db.Find(&user).Error
}
