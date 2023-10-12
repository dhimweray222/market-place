package routes

import (
	"github.com/gin-gonic/gin"
	"market_place.com/controller"
)

func CartRoute(router *gin.Engine)  {

	router.POST("/cart",controller.Authenticate(), controller.AddToCart)
	router.GET("/cart",controller.Authenticate(), controller.GetAllCartFromUser)
	router.DELETE("/cart/:id",controller.Authenticate(), controller.DeleteCart)

}
