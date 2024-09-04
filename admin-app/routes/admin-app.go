package routes

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/admin-app/handlers"
	// "andrenormanlang/tarot-go-htmx/utils"
)


func BackendRegisterRoutes(router *gin.Engine) {

        // API Routes
		router.GET("/cards", handlers.GetCards)
		router.GET("/cards/:id", handlers.GetCardByID)
		router.POST("/cards", handlers.CreateCard)
		router.POST("/bulk-create-cards", handlers.BulkCreateCards)
		// router.POST("/admin/update-images", func(c *gin.Context) {
		// 	utils.UpdateImages()
		// 	c.JSON(200, gin.H{"status": "Images updated"})
		// })
		router.PUT("/cards/:id", handlers.UpdateCard)
		router.PUT("/cards/bulk-update", handlers.BulkUpdateCards)
		router.DELETE("/cards/:id", handlers.DeleteCard)
		router.DELETE("/cards/bulk-delete", handlers.BulkDeleteCards)
}