package main

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/common"
    "andrenormanlang/tarot-go-htmx/app/routes"
    "andrenormanlang/tarot-go-htmx/database"
)

func main() {
    router := gin.Default()
    router.Static("/static", "./static")

    database.ConnectDatabase()

    state := &common.State{
        FullDeck:     []common.Card{}, // Will be filled by the database
        SelectedCards: []common.Card{},
        IsShuffling:  false,
    }

    routes.RegisterRoutes(router, state)

    err := router.Run(":8080") // Listening on port 8080
    if err != nil {
        panic("Server could not start")
    }
}
