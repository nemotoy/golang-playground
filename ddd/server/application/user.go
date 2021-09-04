package application

import (
	"github.com/nemotoy/golang-playground/ddd/server/domain/model"
	"github.com/nemotoy/golang-playground/ddd/server/domain/repository"
)

type UserApplicationService struct {
	userRepo repository.IUserRepository
	// factory
	// domain service
}

func NewUserApplicationService(userRepo repository.IUserRepository) *UserApplicationService {
	return &UserApplicationService{userRepo}
}

func (u *UserApplicationService) GetAll() ([]*model.User, error) {
	users, _ := u.userRepo.GetAll()
	return users, nil
}

func (u *UserApplicationService) Store() (*model.User, error) {
	return &model.User{}, nil
}
