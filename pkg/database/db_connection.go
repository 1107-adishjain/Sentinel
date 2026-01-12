package database

import (
	"log"
	"fmt"
	"gorm.io/driver/postgres"
	"github.com/1107-adishjain/sentinel/pkg/config"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) (*gorm.DB, error){
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s connect_timeout=10 sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPassword)
	db, err:= gorm.Open(postgres.Open((dsn)),&gorm.Config{})

	if err!=nil{
		return nil, err
	}

	return db, nil
}

func CloseDB(db *gorm.DB){
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting database object: %v", err)
		return
	}
	sqlDB.Close()
}