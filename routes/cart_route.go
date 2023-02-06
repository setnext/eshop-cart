package routes

import (
	"eshop-cart-api/controllers"

	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.Engine) {

	router.POST("/cart", controllers.CreateCart())
}
