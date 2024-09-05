package main

import (
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/admin-app/routes"
    "andrenormanlang/tarot-go-htmx/database"
    "andrenormanlang/tarot-go-htmx/common"
    "github.com/gin-contrib/cors"
    "github.com/joho/godotenv"
    "time"
)

func main() {
    // Load .env file if it exists
    godotenv.Load()

    router := gin.Default()

    // Setup CORS middleware with default options
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"https://example.com"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))
    

    router.Static("/static", "./static")
    router.Static("/images", "./images")

    // Initialize and migrate database
    database.ConnectDatabase()
    database.DB.AutoMigrate(&common.Card{})

    // Register backend routes
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
