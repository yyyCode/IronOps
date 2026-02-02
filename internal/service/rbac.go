package service

import (
	"IronOps/internal/database"
	"IronOps/internal/model"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *model.User) error {
	// Hash password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
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

func CheckPassword(user *model.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
