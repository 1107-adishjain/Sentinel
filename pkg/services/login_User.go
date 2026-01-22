package services

import (
	"errors"
	"github.com/1107-adishjain/sentinel/pkg/helper"
	"github.com/1107-adishjain/sentinel/pkg/models"
	"gorm.io/gorm"
)

func LoginUser(db *gorm.DB, email, password string) (string, string, error) {
    var user models.User
    if err := db.Where("email=?", email).First(&user).Error; err != nil {
        return "", "", errors.New("Invalid email")
    }
    if err := helper.VerifyPassword(password, user.PasswordHash); err != nil {
        return "", "", errors.New("invalid password")
    }
    access_token, refresh_token, err := helper.GenerateJWT(user.Id, user.Email)
    if err != nil {
        return "", "", errors.New("failed to generate token")
    }
    return access_token, refresh_token, nil
}	