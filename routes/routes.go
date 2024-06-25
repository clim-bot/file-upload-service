package routes

import (
    "github.com/gin-gonic/gin"

    "github.com/clim-bot/file-upload-service/controllers"
    "github.com/clim-bot/file-upload-service/database"
    "github.com/clim-bot/file-upload-service/middleware"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/upload", middleware.AuthMiddleware(), controllers.UploadFile(database.DB))
    r.GET("/download/:id", middleware.AuthMiddleware(), controllers.DownloadFile(database.DB))
    r.GET("/files", middleware.AuthMiddleware(), controllers.ListFiles(database.DB))
    r.DELETE("/files/:id", middleware.AuthMiddleware(), controllers.DeleteFile(database.DB))

}
