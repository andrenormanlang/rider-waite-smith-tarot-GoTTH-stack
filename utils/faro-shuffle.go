package utils

import "andrenormanlang/tarot-go-htmx/common"

// Faro shuffle function
func FaroShuffle(cards []common.Card) []common.Card {
    // Ensure that the number of cards is even. If not, just return the original deck.
    if len(cards) % 2 != 0 {
        return cards
    }

    // Split the deck into two halves
    mid := len(cards) / 2
    firstHalf := cards[:mid]
    secondHalf := cards[mid:]

    // Interleave the cards from the two halves
    shuffledDeck := make([]common.Card, len(cards))
    for i := 0; i < mid; i++ {
        shuffledDeck[2*i] = firstHalf[i]
        shuffledDeck[2*i+1] = secondHalf[i]
    }

    return shuffledDeck
}