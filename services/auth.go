package services

import (
	"errors"
	"gin-framework-use/internal/model"
	internal "gin-framework-use/internal/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func InitAuthService(db *gorm.DB) *AuthService {
	db.AutoMigrate(&model.User{})
	return &AuthService{
		db: db,
	}
}

func (a *AuthService) CheckUserExist(email *string) bool {
	var user model.User

	if err := a.db.Where("email= ?", email).Find(&user).Error; err != nil {
		return false
	}
	if user.Email != "" {
		return true
	}
	return false
}

func (a *AuthService) Login(email *string, password *string) (*model.User, error) {
	if email == nil {
		return nil, errors.New("email cannot be null")
	}
	if password == nil {
		return nil, errors.New("password cannot be null")

	}
	var user model.User

	if err := a.db.Where("email= ?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	if user.Email == "" {
		return nil, errors.New("no user found")
	}

	if !internal.CheckPasswordHash(*password, user.Password) {
		return nil, errors.New("Password is incorrect")
	}

	return &user, nil
}

func (a *AuthService) Register(email *string, password *string) (*model.User, error) {
	if email == nil {
		return nil, errors.New("email cannot be null")
	}
	if password == nil {
		return nil, errors.New("password cannot be null")

	}

	if a.CheckUserExist(email) {
		return nil, errors.New("user already exist")

	}
	hashedPassword, err := internal.HashPassword(*password)

	var user model.User

	user.Email = *email
	user.Password = hashedPassword

	if err != nil {
		return nil, err

	}

	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
