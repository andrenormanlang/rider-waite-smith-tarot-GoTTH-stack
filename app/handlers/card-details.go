package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"andrenormanlang/tarot-go-htmx/common"
	"andrenormanlang/tarot-go-htmx/database"
	"andrenormanlang/tarot-go-htmx/views"
)

func RevealCardDetail(state *common.State) gin.HandlerFunc {
    return func(c *gin.Context) {
        cardName := c.Query("card")
        var card common.Card

        if err := database.DB.Where("name = ?", cardName).First(&card).Error; err != nil {
            c.String(http.StatusInternalServerError, "Error retrieving card details: %v", err)
            return
        }

        err := views.Modal(card.Name, card.MeaningUp).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering modal template: %v", err)
        }
    }
}
