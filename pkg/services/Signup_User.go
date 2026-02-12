package services

import (
	"errors"
	"github.com/1107-adishjain/sentinel/pkg/helper"
	"github.com/1107-adishjain/sentinel/pkg/models"
	"gorm.io/gorm"
)

func SignupUser(db *gorm.DB, email, password string) error {
	var existingUser models.User
	if err := db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return errors.New("email already in use")
	}
	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user := models.User{
		Email:        email,
		PasswordHash: hashedPassword,
	}
	if err := db.Create(&user).Error; err != nil {
		return errors.New("failed to create user")
	}
	return nil
}
