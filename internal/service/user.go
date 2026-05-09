package service

import (
	"context"
	"github.com/cvrseq/rk-core/internal/models"
)


type UserDataRepository interface {
	CreateUser(ctx context.Context, user *models.VpnConfigModel) error
	GetUserById(ctx context.Context, id int) (models.VpnConfigModel, error)
	GetUsersList(ctx context.Context) ([]models.VpnConfigModel, error)
}

type UserDataService struct { 
	repo UserDataRepository
}

func NewUserDataService(repo UserDataRepository) *UserDataService { 
	return &UserDataService{
		repo: repo,
	}
}

