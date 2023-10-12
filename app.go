package main

import (
	"github.com/gin-gonic/gin"
	"market_place.com/routes"
	"market_place.com/config"
)


func main() {

	router := gin.New()

	config.Connect()
	routes.UserRoute(router)
	routes.AuthRoute(router)
	routes.ProductRoute(router)
	routes.CategoryRoute(router)
	routes.CartRoute(router)
	routes.CheckoutRoute(router)
	router.Run(":8080")

}

