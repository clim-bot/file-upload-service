package controllers

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/clim-bot/file-upload-service/models"
    "github.com/clim-bot/file-upload-service/utils"
)

func UploadFile(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
            return
        }

        fileType := filepath.Ext(file.Filename)
        if !utils.IsValidFileType(fileType) {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type"})
            return
        }

        filePath := fmt.Sprintf("uploads/%s", file.Filename)
        if err := c.SaveUploadedFile(file, filePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
            return
        }

        fileRecord := models.File{
            Name:     file.Filename,
            Path:     filePath,
            Size:     file.Size,
            FileType: fileType,
            Uploader: "uploader", // This should be extracted from JWT token
        }

        db.Create(&fileRecord)
        c.JSON(http.StatusOK, fileRecord)
    }
}

func DownloadFile(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        var file models.File

        if err := db.First(&file, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
            return
        }

        c.File(file.Path)
    }
}

func ListFiles(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var files []models.File
        db.Find(&files)
        c.JSON(http.StatusOK, files)
    }
}

func DeleteFile(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        var file models.File

        if err := db.First(&file, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
            return
        }

        if err := os.Remove(file.Path); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete file"})
            return
        }

        db.Delete(&file)
        c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
    }
}
