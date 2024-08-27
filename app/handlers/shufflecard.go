package handlers

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/common"
    "andrenormanlang/tarot-go-htmx/views"
    "andrenormanlang/tarot-go-htmx/database"
    "net/http"
    "math/rand"
    "time"
)

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
        
        state.FullDeck = shuffleCards(fullDeck)
        state.SelectedCards = []common.Card{} // Reset selectedCards
        
        err := views.Home(state.FullDeck, state.SelectedCards, nil, state.IsShuffling).Render(c.Request.Context(), c.Writer)
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
        
        err := views.Home(state.FullDeck, state.SelectedCards, nil, state.IsShuffling).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
        }
    }
}

func shuffleCards(cards []common.Card) []common.Card {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
    return cards
}
