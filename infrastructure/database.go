package infrastructure

import (
    "fmt"
    "github.com/isd-sgcu/cutu2025-backend/config"
    "github.com/isd-sgcu/cutu2025-backend/domain"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) *gorm.DB {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    // Automatically migrate the schema, creating tables if they don't exist
    err = db.AutoMigrate(&domain.User{}) // Add your domain models here
    if err != nil {
        panic("Failed to auto migrate!")
    }

    return db
}
