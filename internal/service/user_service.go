package service

import (
	mod "rs/internal/models"
)

type UserDataRepository interface {
	CreateUser(user *mod.CfgVpnModel) error
	GetUserById(id int) (mod.CfgVpnModel, error)
	GetUsersList() ([]mod.CfgVpnModel, error)
}

