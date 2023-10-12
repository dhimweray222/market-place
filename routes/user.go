package routes

import (
	"github.com/gin-gonic/gin"
	"market_place.com/controller"
)

func UserRoute(router *gin.Engine)  {
	router.GET("/user",controller.GetUser)
	router.POST("/user",controller.CreateUser)
	router.PUT("/user/:id",controller.UpdateUser)

}
