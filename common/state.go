package common

// Shared state across all handlers
var (
    SelectedCards []Card
    IsShuffling   bool
    FullDeck      []Card
)


type State struct {
    FullDeck     []Card
    SelectedCards []Card
    IsShuffling  bool
    RevealIndex   int 
}
