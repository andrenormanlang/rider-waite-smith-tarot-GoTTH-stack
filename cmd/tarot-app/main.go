package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "math/rand"
    "net/http"
    "time"
    "andrenormanlang/tarot-go-htmx/common"
    "andrenormanlang/tarot-go-htmx/views"
)

var selectedCards []common.Card
var fullDeck []common.Card
var isShuffling bool

func main() {
    router := gin.Default()
    router.Static("/static", "./static")
    fullDeck = generateFullDeck()

    router.GET("/", func(c *gin.Context) {
        err := views.Home(fullDeck, selectedCards, nil, isShuffling).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
        }
    })

    router.GET("/shuffle-cards", func(c *gin.Context) {
        isShuffling = true
        fullDeck = shuffleCards(fullDeck)
        selectedCards = []common.Card{}
        err := views.Home(fullDeck, selectedCards, nil, isShuffling).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
        }
    })

    router.GET("/stop-shuffle", func(c *gin.Context) {
        isShuffling = false
        err := views.Home(fullDeck, selectedCards, nil, isShuffling).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
        }
    })

    router.GET("/select-card", func(c *gin.Context) {
        if !isShuffling && len(selectedCards) < 3 {
            cardName := c.Query("card")
            card := findCardByName(cardName)
            selectedCards = append(selectedCards, card)
        }
        err := views.Home(fullDeck, selectedCards, nil, isShuffling).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
        }
    })

    router.GET("/reveal-meanings", func(c *gin.Context) {
        meanings := generateMeanings(selectedCards)
        err := views.Home(fullDeck, selectedCards, meanings, isShuffling).Render(c.Request.Context(), c.Writer)
        if err != nil {
            c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
        }
    })

    err := router.Run(":8080")
    if err != nil {
        log.Fatalf("Server could not start: %v", err)
    }
}

func shuffleCards(cards []common.Card) []common.Card {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
    return cards
}

func findCardByName(name string) common.Card {
    for _, card := range fullDeck {
        if card.Name == name {
            return card
        }
    }
    return common.Card{}
}

func generateFullDeck() []common.Card {
    return []common.Card{
        {Image: "card1.jpg", Name: "The Fool"},
        {Image: "card2.jpg", Name: "The Magician"},
        {Image: "card3.jpg", Name: "The High Priestess"},
        {Image: "card4.jpg", Name: "The Empress"},
        {Image: "card5.jpg", Name: "The Emperor"},
        {Image: "card6.jpg", Name: "The Hierophant"},
        {Image: "card7.jpg", Name: "The Lovers"},
        {Image: "card8.jpg", Name: "The Chariot"},
        {Image: "card9.jpg", Name: "Strength"},
        {Image: "card10.jpg", Name: "The Hermit"},
        {Image: "card11.jpg", Name: "Wheel of Fortune"},
        {Image: "card12.jpg", Name: "Justice"},
        {Image: "card13.jpg", Name: "The Death"},
        {Image: "card14.jpg", Name: "Temperance"},
        {Image: "card15.jpg", Name: "The Devil"},
        {Image: "card16.jpg", Name: "The Tower"},
        {Image: "card17.jpg", Name: "The Star"},
        {Image: "card18.jpg", Name: "The Moon"},
        {Image: "card19.jpg", Name: "The Sun"},
        {Image: "card20.jpg", Name: "Judgement"},
        {Image: "card21.jpg", Name: "The World"},
    }
}

func generateMeanings(selectedCards []common.Card) []string {
    meanings := []string{}
    for _, card := range selectedCards {
        meanings = append(meanings, "Meaning of "+card.Name)
    }
    return meanings
}