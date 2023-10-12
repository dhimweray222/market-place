package controller

import (
	"github.com/gin-gonic/gin"
	"market_place.com/config"
	"market_place.com/models"
)

type CategoryResponse struct {
    Name  string `json:"name"`
    Description string `json:"description"`
}

func GetCategory(c *gin.Context)  {
	categories := []models.Category{}
	config.DB.Find(&categories)
	c.JSON(200, &categories)
}

func CreateCategory(c *gin.Context)  {
	var category models.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    if err := config.DB.Create(&category).Error; err != nil {
        c.JSON(401, gin.H{"error": err.Error()})
        return
    }
		categoryResponse := CategoryResponse{Name: category.Name, Description: category.Description}
    c.JSON(200, categoryResponse)
}



