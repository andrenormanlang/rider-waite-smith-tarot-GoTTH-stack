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
