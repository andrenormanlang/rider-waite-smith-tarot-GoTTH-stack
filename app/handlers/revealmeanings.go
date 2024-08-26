package handlers

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/common"
    "andrenormanlang/tarot-go-htmx/views"
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
        switch card.Name {
        case "The Fool":
            meanings[i] = "New beginnings, innocence, spontaneity"
        case "The Magician":
            meanings[i] = "Manifestation, resourcefulness, power"
		case "The High Priestess":
			meanings[i] = "Intuition, mystery, subconscious mind"
		case "The Empress":
			meanings[i] = "Fertility, nurturing, abundance"
		case "The Emperor":
			meanings[i] = "Authority, structure, control"
		case "The Hierophant":
			meanings[i] = "Tradition, conformity, spiritual guidance"
		case "The Lovers":
			meanings[i] = "Partnerships, love, choices"
		case "The Chariot":
			meanings[i] = "Determination, willpower, success"
		case "Strength":
			meanings[i] = "Courage, inner strength, compassion"
		case "The Hermit":
			meanings[i] = "Soul-searching, introspection, guidance"
		case "Wheel of Fortune":
			meanings[i] = "Change, cycles, destiny"
		case "Justice":
			meanings[i] = "Fairness, truth, law"
		case "The Hanged Man":
			meanings[i] = "Sacrifice, release, enlightenment"
		case "The Death":
			meanings[i] = "Endings, transformation, new beginnings"
		case "Temperance":
			meanings[i] = "Balance, moderation, harmony"
		case "The Devil":
			meanings[i] = "Materialism, bondage, ignorance"
		case "The Tower":
			meanings[i] = "Sudden change, upheaval, chaos"
		case "The Star":
			meanings[i] = "Hope, inspiration, spirituality"
		case "The Moon":
			meanings[i] = "Illusion, fear, subconscious mind"
		case "The Sun":
			meanings[i] = "Success, joy, vitality"
		case "Judgement":
			meanings[i] = "Judgement, rebirth, inner calling"
		case "The World":
			meanings[i] = "Completion, fulfillment, wholeness"
        default:
            meanings[i] = "Meaning not available for " + card.Name
        }
    }
    return meanings
}
