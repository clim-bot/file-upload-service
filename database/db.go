package database

import (
    "log"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "github.com/clim-bot/file-upload-service/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("file_upload.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    DB.AutoMigrate(&models.File{})
}
