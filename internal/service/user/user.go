package user

import (
	"gopaste/internal/gorm/models"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (s *Service) CreateUser(user *models.User) error {
	return s.DB.Create(user).Error
}

//TODO: Сделать CRUD для юзера

func New(db *gorm.DB) *Service {
	return &Service{DB: db}
}
