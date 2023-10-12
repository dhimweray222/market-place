package controller

import (
	"github.com/gin-gonic/gin"
	"market_place.com/config"
	"market_place.com/models"

)
type ProductWithCategoryData struct {
		Id int `json:"id"`
    Name       string `json:"name"`
    Description      string `json:"description"`
    Category   string `json:"category"`
		Price int `json:"price"`
}
type ProductResponse struct {
		Id int `json:"id"`
    Name  string `json:"name"`
    Description string `json:"description"`
		CategoryId int `json:"categoryId"`
		Price int `json:"price"`
}

func CreateProduct(c *gin.Context)  {
	var product models.Product
	c.ShouldBindJSON(&product);

    if err := config.DB.Create(&product).Error; err != nil {
        c.JSON(401, gin.H{"error": err.Error()})
        return
    }
		productResponse := ProductResponse{Id: product.Id,Name: product.Name, Description: product.Description,CategoryId: product.CategoryId, Price:product.Price}
    c.JSON(200, productResponse)
}

func GetAllProduct(c *gin.Context) {
    var products []models.Product
		var result []ProductWithCategoryData

    if err := config.DB.Preload("Category").Find(&products).Error; err != nil {
        c.JSON(500, gin.H{"error": "Error fetching products"})
        return
    }
		    for _, product := range products {
        result = append(result, ProductWithCategoryData{
						Id : product.Id,
            Name:     product.Name,
            Description:    product.Description,
            Category: product.Category.Name,
						Price : product.Price,
        })
    }
    c.JSON(200, result)
}

func GetProductsByCategoryName(c *gin.Context) {
    categoryName := c.Param("categoryName")
    var products []models.Product

		if err := config.DB.
						Preload("Category").
						Joins("JOIN categories ON products.category_id = categories.id").
						Where("categories.name = ?", categoryName).
						Find(&products).
						Error; err != nil {
				c.JSON(500, gin.H{"error": "Error fetching products"})
				return
		}


    c.JSON(200, products)
}
