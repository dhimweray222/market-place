package controller

import (
	"github.com/gin-gonic/gin"
	"market_place.com/config"
	"market_place.com/models"
)

func Checkout(c *gin.Context) {
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
		var totalPrice int
		result := make([]ProductCategoryData, len(carts))
    for i, cart := range carts {
        result[i] = ProductCategoryData{
						Id: cart.Id,
						User: cart.UserId,
            Product:  cart.Product.Name,
            Category: cart.Product.Category.Name,
        }
				totalPrice += cart.Product.Price
    }

		totalPriceFloat := float64(totalPrice)
		c.Set("totalPrice", totalPriceFloat)
		c.Next()
}
