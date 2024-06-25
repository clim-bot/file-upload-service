package controllers

import (
    "bytes"
    "mime/multipart"
    "net/http"
    "net/http/httptest"
    "os"
    "path/filepath"
    "testing"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "github.com/clim-bot/file-upload-service/models"
)

func TestUploadFile(t *testing.T) {
    db, _ := gorm.Open(sqlite.Open("file_test.db"), &gorm.Config{})
    db.AutoMigrate(&models.File{})

    r := gin.Default()
    r.POST("/upload", UploadFile(db))

    body := new(bytes.Buffer)
    writer := multipart.NewWriter(body)
    file, err := os.Open("testfile.jpg")
    if err != nil {
        t.Fatalf("Failed to open test file: %v", err)
    }
    defer file.Close()

    part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
    if err != nil {
        t.Fatalf("Failed to create form file: %v", err)
    }

    _, err = part.Write([]byte("file content"))
    if err != nil {
        t.Fatalf("Failed to write to part: %v", err)
    }

    writer.Close()

    req, _ := http.NewRequest("POST", "/upload", body)
    req.Header.Set("Content-Type", writer.FormDataContentType())

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("expected %v; got %v", http.StatusOK, w.Code)
    }

    var files []models.File
    db.Find(&files)
    if len(files) == 0 {
        t.Fatalf("expected file to be saved in the database")
    }

    os.Remove(files[0].Path)
}
