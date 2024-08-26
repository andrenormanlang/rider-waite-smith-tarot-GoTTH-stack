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
		err := views.PartialMeanings(meanings).Render(c.Request.Context(), c.Writer)
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
    meanings := make([]string, len(selectedCards))
    for i, card := range selectedCards {
        switch card.Name {
        case "The Fool":
            meanings[i] = "New beginnings, innocence, spontaneity"
        case "The Magician":
            meanings[i] = "Manifestation, resourcefulness, power"
        case "The High Priestess":
            meanings[i] = "Intuition, sacred knowledge, divine feminine"
        case "The Empress":
            meanings[i] = "Femininity, beauty, nature, abundance"
        case "The Emperor":
            meanings[i] = "Authority, establishment, structure, a father figure"
        case "The Hierophant":
            meanings[i] = "Spiritual wisdom, religious beliefs, conformity, tradition"
        case "The Lovers":
            meanings[i] = "Love, harmony, relationships, values alignment, choices"
        case "The Chariot":
            meanings[i] = "Control, willpower, success, action, determination"
        case "Strength":
            meanings[i] = "Strength, courage, persuasion, influence, compassion"
        case "The Hermit":
            meanings[i] = "Soul-searching, introspection, being alone, inner guidance"
        case "Wheel of Fortune":
            meanings[i] = "Good luck, karma, life cycles, destiny, a turning point"
        case "Justice":
            meanings[i] = "Justice, fairness, truth, cause and effect, law"
        case "The Hanged Man":
            meanings[i] = "Surrender, letting go, new perspectives"
        case "Death":
            meanings[i] = "Endings, change, transformation, transition"
        case "Temperance":
            meanings[i] = "Balance, moderation, patience, purpose"
        case "The Devil":
            meanings[i] = "Shadow self, attachment, addiction, restriction, sexuality"
        case "The Tower":
            meanings[i] = "Sudden change, upheaval, chaos, revelation, awakening"
        case "The Star":
            meanings[i] = "Hope, faith, purpose, renewal, spirituality"
        case "The Moon":
            meanings[i] = "Illusion, fear, anxiety, subconscious, intuition"
        case "The Sun":
            meanings[i] = "Positivity, fun, warmth, success, vitality"
        case "Judgement":
            meanings[i] = "Judgement, rebirth, inner calling, absolution"
        case "The World":
            meanings[i] = "Completion, integration, accomplishment, travel"
        default:
            meanings[i] = "Meaning not available for " + card.Name
        }
    }
    return meanings
}