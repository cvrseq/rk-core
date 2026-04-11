package service

import (
	mod "rs/internal/models"
)

type CfgDataRepository interface {
	CreateCfg(cfg mod.CfgVpnModel) error
	CreateOrder(order mod.CfgVpnModel) error
	GetOrderById(orderId int) (mod.CfgVpnModel, error)
	GetOrdersList() ([]mod.CfgVpnModel, error)
	GetCfgByID(id uint) (mod.CfgVpnModel, error)
}

type UseCaseCfg interface {
	// descript usecase contracts
}
