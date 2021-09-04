package repository

import "github.com/nemotoy/golang-playground/ddd/server/domain/model"

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (u *UserRepository) GetAll() []*model.User {
	return nil
}
