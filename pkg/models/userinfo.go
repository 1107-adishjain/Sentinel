package models

import(
	"time"
	"gorm.io/gorm"
)

type User struct{
	Id uint `gorm:"column:id;primaryKey;autoIncrement"`
	Email string `gorm:"column:email;type:varchar(100);uniqueIndex;not null"`
	PasswordHash string `gorm:"column:password_hash;type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}