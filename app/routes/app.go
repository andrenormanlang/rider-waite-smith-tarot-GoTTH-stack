package routes

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/app/handlers"
    "andrenormanlang/tarot-go-htmx/common"
)

func RegisterRoutes(router *gin.Engine, state *common.State) {
    router.GET("/", handlers.Home(state))                // Pass state and get a HandlerFunc
    router.GET("/shuffle-cards", handlers.ShuffleCards(state)) // Same for other routes
    router.GET("/stop-shuffle", handlers.StopShuffle(state))
    router.GET("/select-card", handlers.SelectCard(state))
    router.GET("/reveal-meanings", handlers.RevealMeanings(state))
}
