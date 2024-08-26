package routes

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/admin-app/handlers"
)


func BackendRegisterRoutes(router *gin.Engine) {

        // API Routes
		router.GET("/cards", handlers.GetCards)
		router.GET("/cards/:id", handlers.GetCardByID)
		router.POST("/cards", handlers.CreateCard)
		router.POST("/bulk-create-cards", handlers.BulkCreateCards)
		router.PUT("/cards/:id", handlers.UpdateCard)
		router.DELETE("/cards/:id", handlers.DeleteCard)
}