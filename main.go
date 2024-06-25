package main

import (
    "github.com/gin-gonic/gin"

    "github.com/clim-bot/file-upload-service/config"
    "github.com/clim-bot/file-upload-service/database"
    "github.com/clim-bot/file-upload-service/routes"
)

func main() {
    config.LoadEnv()

    config.LoadEnv()
    database.ConnectDatabase()

    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run()
}
