package models

import(
	"time"
	"gorm.io/gorm"
)

type User struct {
    Id string `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
    Email string `gorm:"column:email;type:varchar(100);uniqueIndex;not null"`
    PasswordHash string `gorm:"column:password_hash;type:varchar(255);not null"`
    CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}