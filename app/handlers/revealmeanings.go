package handlers

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/common"
    "andrenormanlang/tarot-go-htmx/views"
    "andrenormanlang/tarot-go-htmx/database"
    "net/http"
)

func RevealMeanings(state *common.State) gin.HandlerFunc {
    return func(c *gin.Context) {
        meanings := generateMeanings(state.SelectedCards)
        err := views.PartialMeanings(meanings, state.SelectedCards).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
        }
    }
}

func generateMeanings(selectedCards []common.Card) []string {
    meanings := make([]string, len(selectedCards))

    for i, card := range selectedCards {
        var dbCard common.Card
        if err := database.DB.Where("name = ?", card.Name).First(&dbCard).Error; err != nil {
            meanings[i] = "Meaning not available for " + card.Name
        } else {
            meanings[i] = dbCard.MeaningUp // Assuming you're showing the upright meaning; adjust if needed
        }
    }

    return meanings
}
