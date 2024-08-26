package main

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/admin-app/routes"
    "andrenormanlang/tarot-go-htmx/database"
    "andrenormanlang/tarot-go-htmx/models"
)

func main() {
    router := gin.Default()
    router.Static("/static", "./static")

    database.ConnectDatabase()

    // Migrate the schema
    database.DB.AutoMigrate(&models.Card{})


    routes.BackendRegisterRoutes(router)

    err := router.Run(":8081") // Listening on port 8081
    if err != nil {
        panic("Server could not start")
    }
}
