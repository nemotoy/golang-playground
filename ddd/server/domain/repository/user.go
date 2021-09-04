package repository

import "github.com/nemotoy/golang-playground/ddd/server/domain/model"

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (u *UserRepository) GetAll() ([]*model.User, error) {
	return []*model.User{
		{LastName: "aaa", FirstName: "aaa"},
		{LastName: "bbb", FirstName: "bbb"},
	}, nil
}
func (u *UserRepository) Store() (*model.User, error) {
	return nil, nil
}
