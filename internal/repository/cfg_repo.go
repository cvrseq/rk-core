package repository

import (
	mod "rs/internal/models"

	"gorm.io/gorm"
)

type CfgDataRepository interface {
	CreateCfg(cfg mod.CfgVpnModel) error
	CreateOrder(order mod.CfgVpnModel) error
	GetOrderById(orderId int) (mod.CfgVpnModel, error)
	GetOrdersList() ([]mod.CfgVpnModel, error)
	GetCfgByID(id uint) (mod.CfgVpnModel, error)
}

type DBStateCfgs struct {
	db *gorm.DB
}

func NewCfgRepository(db *gorm.DB) CfgDataRepository {
	return &DBStateCfgs{db: db}
}

func (r *DBStateCfgs) CreateCfg(cfg mod.CfgVpnModel) error {
	return r.db.Create(&cfg).Error
}

func (r *DBStateCfgs) CreateOrder(order mod.CfgVpnModel) error {
	return r.db.Create(&order).Error
}

func (r *DBStateCfgs) GetOrderById(orderId int) (mod.CfgVpnModel, error) {
	var cfg mod.CfgVpnModel

	err := r.db.First(&cfg, orderId).Error

	return cfg, err
}

func (r *DBStateCfgs) GetOrdersList() ([]mod.CfgVpnModel, error) {
	var cfg []mod.CfgVpnModel

	return cfg, r.db.Find(&cfg).Error
}

func (r *DBStateCfgs) GetCfgByID(id uint) (mod.CfgVpnModel, error) {
	var cfg mod.CfgVpnModel

	err := r.db.First(&cfg, id).Error

	return cfg, err
}
