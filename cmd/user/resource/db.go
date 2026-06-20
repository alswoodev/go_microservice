package resource

import (
	"fmt"
	"log"
	"userfc/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(cfg *config.Config) *gorm.DB {
	// dsn --> host port user password dbname
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {log.Fatalf("failed connect to DB: %v", err)}
	log.Printf("Connected to DB: %s:%s, name=%s", cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	return db
}