package routes

import (
	"github.com/gin-gonic/gin"
	"market_place.com/controller"
)

func CategoryRoute(router *gin.Engine)  {

	router.POST("/category",controller.Authenticate(), controller.CreateCategory)
	router.GET("/category",controller.GetCategory)
}
