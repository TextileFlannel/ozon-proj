package db

import (
	"log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/TextileFlannel/ozon-proj/graph/model"
)

func InitDB(db *gorm.DB) {
    err := db.AutoMigrate(&model.Post{}, &model.Comment{})
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }
}

func ConnectDB(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}