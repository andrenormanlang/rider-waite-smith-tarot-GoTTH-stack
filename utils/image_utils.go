package utils

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "andrenormanlang/tarot-go-htmx/database"
    "andrenormanlang/tarot-go-htmx/common"  // Use common instead of models
)

// UpdateImages updates the image paths for the cards in the database.
func UpdateImages() {
    imageDir := "./images"

    err := filepath.Walk(imageDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() && isImageFile(info.Name()) {
            cardName, imagePath := extractCardNameFromFilename(info.Name())
            fmt.Printf("Processing file: %s, extracted card name: %s\n", info.Name(), cardName)

            var card common.Card
            if err := database.DB.Where("name = ?", cardName).First(&card).Error; err != nil {
                fmt.Printf("Could not find card %s in the database\n", cardName)
                return nil
            }

            card.Image = imagePath
            database.DB.Save(&card)
            fmt.Printf("Updated card %s with image %s\n", cardName, card.Image)
        }
        return nil
    })

    if err != nil {
        fmt.Printf("Error walking through image directory: %v\n", err)
    }
}


// isImageFile checks if the file is an image based on its extension
func isImageFile(filename string) bool {
    ext := strings.ToLower(filepath.Ext(filename))
    return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
}

// extractCardNameFromFilename extracts the card name from the image filename
func extractCardNameFromFilename(filename string) (string, string) {
    baseName := strings.TrimSuffix(filename, filepath.Ext(filename))

    // Check if the filename matches a Minor Arcana pattern (e.g., "Cups01")
    if len(baseName) > 3 && baseName[len(baseName)-2] >= '0' && baseName[len(baseName)-2] <= '9' {
        suit := baseName[:len(baseName)-2]  // Get the suit (e.g., "Cups")
        number := baseName[len(baseName)-2:] // Get the number (e.g., "01")
        cardName := suit + " " + number
        imagePath := suit + number + ".png"
        return cardName, imagePath
    } else {
        // Handle Major Arcana (e.g., "TheFool" -> "The Fool")
        cardName := strings.ReplaceAll(baseName, "-", " ")
        imagePath := baseName + ".png"
        return cardName, imagePath
    }
}

// GenerateImagePath converts card names to their corresponding image paths
func GenerateImagePath(cardName string) string {
    cardName = strings.ToLower(strings.TrimSpace(cardName))
    
    // Handle Major Arcana (e.g., "The Fool" -> "TheFool.png")
    majorArcana := map[string]string{
        "the fool":        "TheFool.png",
        "the magician":    "TheMagician.png",
        "the high priestess": "TheHighPriestess.png",
        "the empress":     "TheEmpress.png",
        "the emperor":     "TheEmperor.png",
        "the hierophant":  "TheHierophant.png",
        "the lovers":      "TheLovers.png",
        "the chariot":     "TheChariot.png",
        "strength":        "Strength.png",
        "the hermit":      "TheHermit.png",
        "wheel of fortune": "WheelOfFortune.png",
        "justice":         "Justice.png",
        "the hanged man":  "TheHangedMan.png",
        "death":           "Death.png",
        "temperance":      "Temperance.png",
        "the devil":       "TheDevil.png",
        "the tower":       "TheTower.png",
        "the star":        "TheStar.png",
        "the moon":        "TheMoon.png",
        "the sun":         "TheSun.png",
        "judgement":       "Judgement.png",
        "the world":       "TheWorld.png",
    }
    
    if imagePath, exists := majorArcana[cardName]; exists {
        return imagePath
    }

    // Handle Minor Arcana by splitting the name into the rank and suit (e.g., "Four of Swords" -> "swords04.png")
    parts := strings.Split(cardName, " of ")
    if len(parts) == 2 {
        suit := strings.ToLower(parts[1])
        number, err := convertToNumber(parts[0])
        if err == nil {
            return fmt.Sprintf("%s%02d.png", suit, number)
        }
    }

    // If no match found, return an empty string or some default image path
    return ""
}


// convertToNumber converts the card rank (e.g., "Four") to its corresponding number
func convertToNumber(rank string) (int, error) {
    ranks := map[string]int{
        "ace":    1,
        "two":    2,
        "three":  3,
        "four":   4,
        "five":   5,
        "six":    6,
        "seven":  7,
        "eight":  8,
        "nine":   9,
        "ten":    10,
        "page":   11,
        "knight": 12,
        "queen":  13,
        "king":   14,
    }

    if number, found := ranks[strings.ToLower(rank)]; found {
        return number, nil
    }
    
    return 0, fmt.Errorf("invalid rank: %s", rank)
}
