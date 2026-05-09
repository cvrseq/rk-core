package repository

import (
	"context"

	"github.com/cvrseq/rk-core/internal/models"

	"gorm.io/gorm"
)

type VpnConfigRepository struct {
	db *gorm.DB
}

func NewVpnConfigRepository(db *gorm.DB) *VpnConfigRepository {
	return &VpnConfigRepository{db: db}
}

func (r *VpnConfigRepository) Create(ctx context.Context, cfg models.VpnConfigModel) error {
	return r.db.Create(&cfg).Error
}

func (r *VpnConfigRepository) List(ctx context.Context) ([]models.VpnConfigModel, error) {
	var cfg []models.VpnConfigModel

	return cfg, r.db.Find(&cfg).Error
}

func (r *VpnConfigRepository) GetByID(ctx context.Context, id uint) (models.VpnConfigModel, error) {
	var cfg models.VpnConfigModel

	err := r.db.First(&cfg, id).Error

	return cfg, err
}
