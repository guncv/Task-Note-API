package database

import (
	"fmt"
	"log"

	"github.com/guncv/tech-exam-software-engineering/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DbName, cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database object from GORM: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	return db
}
