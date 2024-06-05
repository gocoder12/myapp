package repository

import "myapp/internal/model"

type UserRepository interface {
	FetchAllUsers() ([]model.User, error)
	FetchUserByID(id string) (model.User, error)
	CreateUser(user model.User) error
	UpdateUser(id string, user model.User) error
	DeleteUser(id string) error
}
