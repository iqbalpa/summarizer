package service

import (
	"summarizer/internal/model"
	"summarizer/internal/repo"
)

type IUserService interface {
	CreateUser(name, username string) (model.User, error)
}

type UserService struct {
	ur repo.IUserRepository
}

func NewUserService(ur repo.IUserRepository) *UserService {
	return &UserService{
		ur: ur,
	}
}

func (us *UserService) CreateUser(name, username string) (model.User, error) {
	user, err := us.ur.CreateUser(name, username)
	return user, err
}
