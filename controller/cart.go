package controller

import (
	"github.com/gin-gonic/gin"
	"market_place.com/config"
	"market_place.com/models"
)

func AddToCart(c *gin.Context) {
    iduser, exists := c.Get("userID")
    if !exists {
        c.JSON(401, gin.H{"error": "User data not found"})
        return
    }

    var request struct {
        ProductId int `json:"ProductId"`
    }

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    cart := models.Cart{
        UserId:    iduser.(int),
        ProductId: request.ProductId,
    }

    if err := config.DB.Create(&cart).Error; err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Item added to cart", "cart": cart})
}

func GetAllCartFromUser(c *gin.Context) {
    var carts []models.Cart
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(401, gin.H{"error": "User data not found"})
        return
    }
		if err := config.DB.
        Preload("Product.Category").
        Where("user_id = ?", userID).
        Find(&carts).
				Error; err != nil {
        c.JSON(500, gin.H{"error": "Error fetching products"})
        return
    }
		type ProductCategoryData struct {
			Id int `json:"Id"`
			User int
			Product  string `json:"product"`
			Category string `json:"category"`
		}
		result := make([]ProductCategoryData, len(carts))
    for i, cart := range carts {
        result[i] = ProductCategoryData{
						Id: cart.Id,
						User: cart.UserId,
            Product:  cart.Product.Name,
            Category: cart.Product.Category.Name,
        }
    }
    c.JSON(200, result)
}



func DeleteCart(c *gin.Context)  {
	userID, exists := c.Get("userID")
	productID := c.Param("id")
	if !exists {
			c.JSON(401, gin.H{"error": "User data not found"})
			return
	}
	var cart models.Cart
	if err := config.DB.Where("user_id = ? AND id = ?", userID, productID).First(&cart).Error; err != nil {
		c.JSON(404,"data doesn't exist")
	} else {
			if err := config.DB.Delete(&cart).Error; err != nil {
					c.JSON(500, "Internal Server Error")
			} else {
					c.JSON(200, "Successfully Deleted Data")
			}
	}
}
