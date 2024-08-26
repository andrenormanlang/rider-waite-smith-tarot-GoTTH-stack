package main

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/common"
    "andrenormanlang/tarot-go-htmx/app/routes"
)

func main() {
    router := gin.Default()
    router.Static("/static", "./static")

    state := &common.State{
        FullDeck:     common.GenerateFullDeck(),
        SelectedCards: []common.Card{},
        IsShuffling:  false,
    }

    routes.RegisterRoutes(router, state)

    err := router.Run(":8080") // Listening on port 8080
    if err != nil {
        panic("Server could not start")
    }
}
