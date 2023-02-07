package routes

import (
	"eshop-cart-api/controllers"

	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.Engine) {

	router.POST("/cart", controllers.CreateOrUpdateCart())

	router.GET("/cart/:uid", controllers.GetCart())

	router.DELETE("/cart/:uid", controllers.DeleteCart())

	router.GET("/carts", controllers.GetAllCarts())
}
