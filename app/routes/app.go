package routes

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/app/handlers"
    "andrenormanlang/tarot-go-htmx/common"
)

func RegisterRoutes(router *gin.Engine, state *common.State) {
    router.GET("/", handlers.Home(state))                
    router.GET("/shuffle-cards", handlers.ShuffleCards(state)) 
    router.GET("/stop-shuffle", handlers.StopShuffle(state))
    router.GET("/select-card", handlers.SelectCard(state))
    router.GET("/reveal-meanings", handlers.RevealMeanings(state))
    router.GET("/card-detail", handlers.RevealCardDetail(state))
    router.GET("/reset-reading", handlers.ResetReadingHandler(state)) // Correct the path to your handler
}
