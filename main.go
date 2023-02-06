package main

import (
	"eshop-cart-api/configs"
	"eshop-cart-api/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

//add this

func main() {
	router := gin.Default()

	configs.ConnectDB()

	//routes
	routes.CartRoute(router) //add this

	router.GET("/", func(c *gin.Context) {
		fmt.Print("Call Received")
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})

	router.Run("localhost:6000")
}
