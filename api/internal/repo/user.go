package repo

import (
	"fmt"
	"summarizer/internal/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(name, username, encryptedPass string) (model.User, error)
	GetUser(username string) (model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(name, username, encryptedPass string) (model.User, error) {
	user := model.User{
		Name:     name,
		Username: username,
		Password: encryptedPass,
	}
	res := ur.db.Create(&user)
	if res.Error != nil {
		return model.User{}, fmt.Errorf("failed to create a new user")
	}
	return user, nil
}

func (ur *UserRepository) GetUser(username string) (model.User, error) {
	var user model.User
	result := ur.db.Where("username = ?", username).Find(&user)
	if result.Error != nil {
		return model.User{}, fmt.Errorf("user with username %s is not found", username)
	}
	return user, nil
}
