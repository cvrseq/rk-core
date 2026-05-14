package service

import (
	"context"
	"errors"

	"github.com/cvrseq/rk-core/internal/handlers"
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

func (s *VpnConfigDataService) GenerateConfig(ctx context.Context, input handlers.VpnConfigRequest) (handlers.VpnConfigResponse, error) {
	envCfg := "some generated config" // hardcode

	if input.UserID == 0 {
		return handlers.VpnConfigResponse{}, errors.New("user_id is required")
	}

	if input.OrderID == 0 {
		return handlers.VpnConfigResponse{}, errors.New("order_id is required")
	}

	if input.Region == "" {
		return handlers.VpnConfigResponse{}, errors.New("region is required")
	}

	config := models.VpnConfigModel{
		UserID:  input.UserID,
		OrderID: input.OrderID,
		Region:  input.Region,
	}

	if err := s.repo.Create(ctx, config); err != nil {
		return handlers.VpnConfigResponse{}, err
	}

	return handlers.VpnConfigResponse{
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
