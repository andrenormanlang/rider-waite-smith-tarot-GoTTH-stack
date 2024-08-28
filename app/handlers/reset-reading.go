package handlers

import (
    "github.com/gin-gonic/gin"
    "log"
    "andrenormanlang/tarot-go-htmx/common"
    "andrenormanlang/tarot-go-htmx/views"
    "net/http"
)

func ResetReadingHandler(state *common.State) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Clear selected cards
        state.SelectedCards = []common.Card{}
        state.IsShuffling = false  // Ensure that shuffling is stopped

        // Log that we've entered this handler
        log.Println("Reset reading handler called")

         // Add a custom header to indicate a full reload should occur
         c.Header("HX-Refresh", "true")

        // Re-render the entire page
        err := views.Home(state.FullDeck, state.SelectedCards, nil, state.IsShuffling).Render(c.Request.Context(), c.Writer)
        if err != nil {
            log.Printf("Error rendering home template: %v", err)
            c.String(http.StatusInternalServerError, "Internal Server Error")
        }
    }
}
