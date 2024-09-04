package handlers

import (
    "github.com/gin-gonic/gin"
    "andrenormanlang/tarot-go-htmx/database"
    "andrenormanlang/tarot-go-htmx/common"
    "net/http"
)

func GetCards(c *gin.Context) {
    var cards []common.Card
    database.DB.Find(&cards)
    c.JSON(http.StatusOK, cards)
}

func GetCardByID(c *gin.Context) {
    id := c.Param("id")
    var card common.Card
    if err := database.DB.First(&card, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
        return
    }
    c.JSON(http.StatusOK, card)
}

func CreateCard(c *gin.Context) {
    var input common.Card
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&input)
    c.JSON(http.StatusOK, input)
}

func BulkCreateCards(c *gin.Context) {
    var cards []common.Card
    if err := c.ShouldBindJSON(&cards); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := database.DB.Create(&cards)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Cards created successfully", "count": result.RowsAffected})
}

func UpdateCard(c *gin.Context) {
    id := c.Param("id")
    var card common.Card
    if err := database.DB.First(&card, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
        return
    }

    var input common.Card
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Model(&card).Updates(input)
    c.JSON(http.StatusOK, card)
}

func BulkUpdateCards(c *gin.Context) {
    var cards []common.Card
    if err := c.ShouldBindJSON(&cards); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for _, card := range cards {
        var existingCard common.Card
        if err := database.DB.First(&existingCard, card.ID).Error; err != nil {
            // If the card is not found, return an error and skip the update
            c.JSON(http.StatusNotFound, gin.H{"error": "Card not found for ID " + string(card.ID)})
            continue
        }

        // Update the card record with new data
        if err := database.DB.Model(&existingCard).Updates(card).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update card with ID " + string(card.ID)})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Cards updated successfully"})
}

func DeleteCard(c *gin.Context) {
    id := c.Param("id")
    var card common.Card
    if err := database.DB.First(&card, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
        return
    }

    database.DB.Delete(&card)
    c.JSON(http.StatusOK, gin.H{"data": "Card deleted"})
}

type BulkDeleteRequest struct {
    IDs []int `json:"ids"`
}

func BulkDeleteCards(c *gin.Context) {
    var req BulkDeleteRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    if len(req.IDs) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No IDs provided"})
        return
    }

    result := database.DB.Delete(&common.Card{}, req.IDs)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Cards deleted successfully", "count": result.RowsAffected})
}


