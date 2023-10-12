package controller

import (
	"github.com/gin-gonic/gin"
	"market_place.com/config"
	"market_place.com/models"
)

type UserResponse struct {
    Id       int   `json:"id"`
    Name string `json:"username"`
		Email string `json:"email"`
}

func GetUser(c *gin.Context)  {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func CreateUser(c *gin.Context)  {
	var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(401, gin.H{"error": err.Error()})
        return
    }
		 userResponse := UserResponse{
        Id:       user.Id,
        Name: user.Name,
				Email: user.Email,
    }
    c.JSON(200, &userResponse)
}

func UpdateUser(c *gin.Context)  {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.ShouldBindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)

}
