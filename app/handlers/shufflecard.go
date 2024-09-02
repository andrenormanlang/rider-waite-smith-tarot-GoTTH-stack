package handlers

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/common"
    "andrenormanlang/tarot-go-htmx/views"
    "andrenormanlang/tarot-go-htmx/database"
    "andrenormanlang/tarot-go-htmx/utils"
    "net/http"
)

// ShuffleCards applies the Faro shuffle to the deck
func ShuffleCards(state *common.State) gin.HandlerFunc {
    return func(c *gin.Context) {
        state.IsShuffling = true
        
        // Fetch the full deck from the database
        var fullDeck []common.Card
        database.DB.Find(&fullDeck)
        
        // Set the image to CardBacks for all cards
        for i := range fullDeck {
            fullDeck[i].Image = "CardBacks.png"
        }
        
        // Apply the Faro shuffle instead of the random shuffle
        // state.FullDeck = utils.PerformMultipleFaroShuffles(fullDeck, 3)
        state.FullDeck = utils.PerformMultipleRiffleShuffles(fullDeck, 3) 
        state.SelectedCards = []common.Card{} // Reset selectedCards
        
        err := views.Home(state.FullDeck, state.SelectedCards, nil, state.IsShuffling, state.RevealIndex).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
        }
    }
}

func StopShuffle(state *common.State) gin.HandlerFunc {
    return func(c *gin.Context) {
        state.IsShuffling = false
        state.SelectedCards = []common.Card{} // Reset selectedCards
        
        // Reset images back to CardBacks
        for i := range state.FullDeck {
            state.FullDeck[i].Image = "CardBacks.png"
        }
        
        err := views.Home(state.FullDeck, state.SelectedCards, nil, state.IsShuffling, state.RevealIndex).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
        }
    }
}

// Faro shuffle function
// func faroShuffle(cards []common.Card) []common.Card {
//     // Ensure that the number of cards is even. If not, just return the original deck.
//     if len(cards) % 2 != 0 {
//         return cards
//     }

//     // Split the deck into two halves
//     mid := len(cards) / 2
//     firstHalf := cards[:mid]
//     secondHalf := cards[mid:]

//     // Interleave the cards from the two halves
//     shuffledDeck := make([]common.Card, len(cards))
//     for i := 0; i < mid; i++ {
//         shuffledDeck[2*i] = firstHalf[i]
//         shuffledDeck[2*i+1] = secondHalf[i]
//     }

//     return shuffledDeck
// }
