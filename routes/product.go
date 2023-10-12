package routes

import (
	"github.com/gin-gonic/gin"
	"market_place.com/controller"
)

func ProductRoute(router *gin.Engine)  {
	router.GET("/product",controller.GetAllProduct)
	router.GET("/product/:categoryName",controller.GetProductsByCategoryName)
	router.POST("/product",controller.CreateProduct)

}
