package models

import (
	"time"
	"gorm.io/gorm"
)

type RateLimitEvent struct {
	ID uint `gorm:"column:id;primaryKey;autoIncrement"`
	UserID string `gorm:"column:user_id;type:varchar(64);index;not null"`
	Endpoint string `gorm:"column:endpoint;type:varchar(255);not null"`
	Method string `gorm:"column:method;type:varchar(10);not null"`
	IP string `gorm:"column:ip;type:varchar(45);not null"`
	Reason string `gorm:"column:reason;type:varchar(50);not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}


func MigrateRateLimitEvent(db *gorm.DB) error {
	return db.AutoMigrate(&RateLimitEvent{})
}