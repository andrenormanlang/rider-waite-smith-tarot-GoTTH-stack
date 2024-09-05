package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/admin-app/routes"
    "andrenormanlang/tarot-go-htmx/database"
    "andrenormanlang/tarot-go-htmx/common"
    "github.com/joho/godotenv"
     "github.com/gin-contrib/cors"
    "time"
)

func main() {

    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    router := gin.Default()
    router.Static("/static", "./static")
    router.Static("/images", "./images")  // Serve images directory

     // Configure CORS middleware
     router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"https://rider-waite-smith-tarot.onrender.com/"}, // Adjust this to your frontend's URL
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

    database.ConnectDatabase()

    // Migrate the schema to ensure the database structure is up-to-date
    database.DB.AutoMigrate(&common.Card{})

    // Register all admin routes
    routes.BackendRegisterRoutes(router)

    // Start the server on port 8081
    err = router.Run(":8081")
    if err != nil {
        panic("Server could not start")
    }
}