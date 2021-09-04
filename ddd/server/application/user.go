package application

import "github.com/nemotoy/golang-playground/ddd/server/domain/model"

type UserApplicationService struct {
	// repository interface
	// factory
	// domain service
}

func NewUserApplicationService() *UserApplicationService {
	return &UserApplicationService{}
}

func (u *UserApplicationService) GetAll() []*model.User {
	return []*model.User{}
}
