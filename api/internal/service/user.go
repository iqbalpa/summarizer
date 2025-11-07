package service

import (
	"summarizer/internal/model"
	"summarizer/internal/repo"
	"summarizer/internal/utils"
)

type IUserService interface {
	CreateUser(name, username, encryptedPass string) (model.User, error)
}

type UserService struct {
	ur repo.IUserRepository
}

func NewUserService(ur repo.IUserRepository) *UserService {
	return &UserService{
		ur: ur,
	}
}

func (us *UserService) CreateUser(name, username, password string) (model.User, error) {
	encryptedPass, err := utils.HashPassword(password)
	if err != nil {
		return model.User{}, err
	}
	user, err := us.ur.CreateUser(name, username, encryptedPass)
	return user, err
}
