package handlers

import (
	"andrenormanlang/tarot-go-htmx/common"
	"andrenormanlang/tarot-go-htmx/views"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func SelectCard(state *common.State) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !state.IsShuffling && len(state.SelectedCards) < 3 {
			cardName, _ := url.QueryUnescape(c.Query("card"))
			card := findCardByName(cardName, state.FullDeck)

			// Check if the card is already selected
			alreadySelected := false
			for _, selectedCard := range state.SelectedCards {
				if selectedCard.Name == card.Name {
					alreadySelected = true
					break
				}
			}

			if !alreadySelected {
				// Now using the image field directly from the card JSON object
				fmt.Printf("Selected card: %s, Image path: %s\n", card.Name, card.Image)

				state.SelectedCards = append(state.SelectedCards, card)
			}
		}

		err := views.Home(state.FullDeck, state.SelectedCards, nil, state.IsShuffling, state.RevealIndex).Render(c.Request.Context(), c.Writer)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
			return
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
