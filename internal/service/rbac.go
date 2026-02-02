package service

import (
	"IronOps/internal/database"
	"IronOps/internal/model"
)

func CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

func ListUsers() ([]model.User, error) {
	var users []model.User
	err := database.DB.Find(&users).Error
	return users, err
}
