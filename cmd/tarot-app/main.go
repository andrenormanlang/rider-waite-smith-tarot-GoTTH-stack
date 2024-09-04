package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/common"
    "andrenormanlang/tarot-go-htmx/app/routes"
    "github.com/joho/godotenv"
    "andrenormanlang/tarot-go-htmx/database"
)

func main() {

    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    router := gin.Default()
    router.Static("/static", "./static")
    router.Static("/images", "./images")  // Serve images directory

    database.ConnectDatabase()

    var fullDeck []common.Card
    database.DB.Find(&fullDeck)  // Populate FullDeck from the database

    state := &common.State{
        FullDeck:     fullDeck,   // Set FullDeck from the database
        SelectedCards: []common.Card{},
        IsShuffling:  false,
        RevealIndex:   0,
    }

    // Register all application routes
    routes.RegisterRoutes(router, state)

    // Start the server on port 8080
    err = router.Run(":8080")  // Change := to =
    if err != nil {
        panic("Server could not start")
    }
}
