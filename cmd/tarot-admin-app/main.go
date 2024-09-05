package main

import (
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/admin-app/routes"
    "andrenormanlang/tarot-go-htmx/database"
    "andrenormanlang/tarot-go-htmx/common"
    "github.com/joho/godotenv"
)

func main() {
    // Load .env file if it exists
    godotenv.Load()

    router := gin.Default()
    router.Static("/static", "./static")
    router.Static("/images", "./images")

    database.ConnectDatabase()
    database.DB.AutoMigrate(&common.Card{})

    routes.BackendRegisterRoutes(router)

    // Get port from environment variable
    port := os.Getenv("PORT")
    if port == "" {
        port = "8081" // Default to 8081 if PORT is not set
    }

    // Start the server
    err := router.Run(":" + port)
    if err != nil {
        log.Fatalf("Server could not start: %v", err)
    }
}