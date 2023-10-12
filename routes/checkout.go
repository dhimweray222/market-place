package routes

import (
	"github.com/gin-gonic/gin"
	"market_place.com/controller"
)

func CheckoutRoute(router *gin.Engine)  {
	router.POST("/checkout",controller.Authenticate(), controller.Checkout, controller.Payment)
}
