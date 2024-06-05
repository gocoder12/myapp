package usecase

import (
	"myapp/internal/model"
	"myapp/internal/repository"
	"myapp/pkg/util"
)

type UserUsecase interface {
	FetchAllUsers() ([]model.User, error)
	FetchUserByID(id string) (model.User, error)
	CreateUser(user model.User) error
	UpdateUser(id string, user model.User) error
	DeleteUser(id string) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase() UserUsecase {
	db := util.ConnectDB()
	userRepo := repository.NewUserRepository(db)
	return &userUsecase{userRepo}
}

func (u *userUsecase) FetchAllUsers() ([]model.User, error) {
	return u.userRepository.FetchAllUsers()
}

func (u *userUsecase) FetchUserByID(id string) (model.User, error) {
	return u.userRepository.FetchUserByID(id)
}

func (u *userUsecase) CreateUser(user model.User) error {
	return u.userRepository.CreateUser(user)
}

func (u *userUsecase) UpdateUser(id string, user model.User) error {
	return u.userRepository.UpdateUser(id, user)
}

func (u *userUsecase) DeleteUser(id string) error {
	return u.userRepository.DeleteUser(id)
}
