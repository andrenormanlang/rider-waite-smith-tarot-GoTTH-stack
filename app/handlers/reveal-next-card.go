package handlers

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/common"
    "andrenormanlang/tarot-go-htmx/views"
    "net/http"
)

func RevealNextCard(state *common.State) gin.HandlerFunc {
    return func(c *gin.Context) {
        if state.RevealIndex < len(state.SelectedCards) {
            state.RevealIndex++
        }

        err := views.Home(state.FullDeck, state.SelectedCards, nil, state.IsShuffling, state.RevealIndex).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
            return
        }
    }
}
