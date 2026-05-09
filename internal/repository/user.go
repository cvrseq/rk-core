package repository

import (
	"github.com/cvrseq/rk-core/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.UserModel) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) GetUserById(id int) (models.UserModel, error) {
	var user models.UserModel

	err := r.db.First(&user, id).Error

	return user, err
}

func (r *UserRepository) GetUsersList() ([]models.UserModel, error) {
	var user []models.UserModel

	return user, r.db.Find(&user).Error
}
