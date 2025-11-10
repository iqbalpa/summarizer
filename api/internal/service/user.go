package service

import (
	"summarizer/internal/model"
	"summarizer/internal/repo"
	"summarizer/internal/utils"
)

type IUserService interface {
	CreateUser(name, username, password string) (model.User, error)
	Login(username, password string) (string, error)
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
	user.Password = ""
	return user, err
}

func (us *UserService) Login(username, password string) (string, error) {
	user, err := us.ur.GetUser(username)
	if err != nil {
		return "failed to login", err
	}
	// check the password
	_, err = utils.ComparePassword(password, user.Password)
	if err != nil {
		return "", err
	}
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}
