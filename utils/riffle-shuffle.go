package utils

import (
    "math/rand"
    "andrenormanlang/tarot-go-htmx/common"
)

// RiffleShuffle performs a single riffle shuffle on a deck of cards
func RiffleShuffle(cards []common.Card) []common.Card {
    // Split the deck into two halves
    mid := len(cards) / 2

    firstHalf := cards[:mid]
    secondHalf := cards[mid:]

    shuffledDeck := make([]common.Card, 0, len(cards))

    // Randomly interleave cards from the two halves
    for len(firstHalf) > 0 || len(secondHalf) > 0 {
        if len(firstHalf) > 0 && (len(secondHalf) == 0 || rand.Intn(2) == 0) {
            shuffledDeck = append(shuffledDeck, firstHalf[0])
            firstHalf = firstHalf[1:]
        }
        if len(secondHalf) > 0 && (len(firstHalf) == 0 || rand.Intn(2) == 1) {
            shuffledDeck = append(shuffledDeck, secondHalf[0])
            secondHalf = secondHalf[1:]
        }
    }

    return shuffledDeck
}

// PerformMultipleRiffleShuffles shuffles the deck multiple times using the riffle shuffle technique
func PerformMultipleRiffleShuffles(cards []common.Card, times int) []common.Card {
    shuffledDeck := cards
    for i := 0; i < times; i++ {
        shuffledDeck = RiffleShuffle(shuffledDeck)
    }
    return shuffledDeck
}
