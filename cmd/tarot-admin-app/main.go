package main

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/admin-app/routes"
    "andrenormanlang/tarot-go-htmx/database"
    "andrenormanlang/tarot-go-htmx/common"
)

func main() {
    router := gin.Default()
    router.Static("/static", "./static")
    router.Static("/images", "./images")  // Serve images directory

    database.ConnectDatabase()

    // Migrate the schema to ensure the database structure is up-to-date
    database.DB.AutoMigrate(&common.Card{})

    // Register all admin routes
    routes.BackendRegisterRoutes(router)

    // Start the server on port 8081
    err := router.Run(":8081")
    if err != nil {
        panic("Server could not start")
    }
}
