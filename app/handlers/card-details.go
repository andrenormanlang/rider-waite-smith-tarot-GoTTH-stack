package handlers

import (
	"github.com/gin-gonic/gin"
	"andrenormanlang/tarot-go-htmx/common"
	"andrenormanlang/tarot-go-htmx/views"
	"andrenormanlang/tarot-go-htmx/database"
	"net/http"
)

func RevealCardDetail(state *common.State) gin.HandlerFunc {
    return func(c *gin.Context) {
        cardName := c.Query("card")
        var card common.Card

        if err := database.DB.Where("name = ?", cardName).First(&card).Error; err != nil {
            c.String(http.StatusInternalServerError, "Error retrieving card details: %v", err)
            return
        }

        err := views.CardDetailModal(card.Name, card.Description, card.MeaningUp, card.MeaningRev ).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering modal template: %v", err)
        }
    }
}
