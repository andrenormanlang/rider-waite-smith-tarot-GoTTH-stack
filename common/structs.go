package common

import "gorm.io/gorm"

type Card struct {
    gorm.Model
    Name        string `json:"name"`        // Full name of the card (e.g., "The Fool")
    NameShort   string `json:"name_short"`  // Short name or identifier (e.g., "Fool")
    Type        string `json:"type"`        // Card type (e.g., "Major Arcana")
    Value       string `json:"value"`       // Value or rank of the card (e.g., "0")
    ValueInt    int    `json:"value_int"`   // Numerical value of the card (e.g., 0 for The Fool)
    MeaningUp   string `json:"meaning_up"`  // Upright meaning of the card
    MeaningRev  string `json:"meaning_rev"` // Reversed meaning of the card
    Description string `json:"desc"`        // Full description of the card
    Image       string `json:"image"`       // Filename or path to the card's image
}
