package routes

import (
	"github.com/gin-gonic/gin"
	"market_place.com/controller"
)

func AuthRoute(router *gin.Engine)  {

	router.POST("/login",controller.GenerateToken)
}
