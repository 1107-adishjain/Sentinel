package controllers

import (
	"net/http"
	svc "github.com/1107-adishjain/sentinel/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Signup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email     string `json:"email" validate:"required,email"`
			Password  string `json:"password" validate:"required,min=8"`
		}
		err:= c.BindJSON(&req); if err!=nil{
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}
		if req.Email=="" || req.Password==""{
			c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
			return
		}	
		if err := validator.New().Struct(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "fields are not validated"})
			return
		}
		err = svc.SignupUser(db,req.Email, req.Password)
		if err != nil {
			if err.Error() == "email already in use" {
				c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully"})
	}
}
func Login(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Email    string `json:"email" validate:"required,email"`
            Password string `json:"password" validate:"required,min=8"`
        }
        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
            return
        }
        if err := validator.New().Struct(req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "fields are not validated"})
            return
        }
        accessToken, refreshToken, err := svc.LoginUser(db, req.Email, req.Password)
        if err != nil {
            if err.Error() == "Invalid email" {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
                return
            }
            if err.Error() == "invalid password" {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
                return
            }
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
		c.SetCookie(
			"refresh_token",
			refreshToken,
			7*24*60*60,
			"/",
			"",
			true,
			true,
		)
        c.JSON(http.StatusOK, gin.H{
            "access_token": accessToken,
        })
    }
}