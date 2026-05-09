package service

import (
	"context"

	"github.com/cvrseq/rk-core/internal/models"
)

type VpnConfigDataRepository interface {
	Create(ctx context.Context, cfg models.VpnConfigModel) error
	List(ctx context.Context) ([]models.VpnConfigModel, error)
	GetByID(ctx context.Context, id uint) (models.VpnConfigModel, error)
}

type VpnConfigDataService struct {
	repo VpnConfigDataRepository
}

func NewVpnConfigDataService(repo VpnConfigDataRepository) *VpnConfigDataService {
	return &VpnConfigDataService{
		repo: repo,
	}
}

type VpnConfigRequest struct {
	UserID  uint   `json:"user_id"`
	OrderID uint   `json:"order_id"`
	Region  string `json:"region"`
}

type VpnConfigResponse struct {
	UserID  uint   `json:"user_id"`
	OrderID uint   `json:"order_id"`
	Region  string `json:"region"`
	Cfg     string `json:"config"`
}

func (s *VpnConfigDataService) GenerateConfig(ctx context.Context, input VpnConfigRequest) (VpnConfigResponse, error) {
	envCfg := "some generated config"

	config := models.VpnConfigModel{
		UserID:  input.UserID,
		OrderID: input.OrderID,
	}

	if err := s.repo.Create(ctx, config); err != nil {
		return VpnConfigResponse{}, err
	}

	return VpnConfigResponse{
		UserID:  input.UserID,
		OrderID: input.OrderID,
		Region:  input.Region,
		Cfg:     envCfg,
	}, nil
}

func (s *VpnConfigDataService) GetConfig(ctx context.Context) ([]models.VpnConfigModel, error) {
	configs, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return configs, nil
}
