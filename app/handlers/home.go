package handlers

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/common"
    "andrenormanlang/tarot-go-htmx/views"
    "net/http"
)

func Home(state *common.State) gin.HandlerFunc {
    return func(c *gin.Context) {
        err := views.Home(state.FullDeck, state.SelectedCards, nil, state.IsShuffling).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
        }
    }
}
