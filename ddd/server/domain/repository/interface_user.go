package repository

import "github.com/nemotoy/golang-playground/ddd/server/domain/model"

type IUserRepository interface {
	GetAll() ([]*model.User, error)
	Store() (*model.User, error)
}
