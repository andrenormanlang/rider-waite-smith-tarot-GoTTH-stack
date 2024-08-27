package utils

import (
	"andrenormanlang/tarot-go-htmx/common"

)
func PerformMultipleFaroShuffles(cards []common.Card, times int) []common.Card {
    shuffledDeck := cards
    for i := 0; i < times; i++ {
        shuffledDeck = FaroShuffle(shuffledDeck)
    }
    return shuffledDeck
}
