package handlers

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/common"
    "andrenormanlang/tarot-go-htmx/views"
    "net/http"
)

func SelectCard(state *common.State) gin.HandlerFunc {
    return func(c *gin.Context) {
        if !state.IsShuffling && len(state.SelectedCards) < 3 {
            cardName := c.Query("card")
            card := findCardByName(cardName, state.FullDeck)

            // Check if the card is already selected
            alreadySelected := false
            for _, selectedCard := range state.SelectedCards {
                if selectedCard.Name == card.Name {
                    alreadySelected = true
                    break
                }
            }

            // Append the card if it's not already selected
            if !alreadySelected {
                state.SelectedCards = append(state.SelectedCards, card)
            }
        }
        err := views.Home(state.FullDeck, state.SelectedCards, nil, state.IsShuffling).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
        }
    }
}

func findCardByName(name string, fullDeck []common.Card) common.Card {
    for _, card := range fullDeck {
        if card.Name == name {
            return card
        }
    }
    return common.Card{}
}

