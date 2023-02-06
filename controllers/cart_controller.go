package controllers

import (
	"context"
	"eshop-cart-api/configs"
	"eshop-cart-api/models"
	"eshop-cart-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection = configs.GetCollection(configs.DB, "carts")
var validate = validator.New()

func CreateCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var cart models.Cart
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, responses.CartResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&cart); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.CartResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newCart := models.Cart{
			Id:        primitive.NewObjectID(),
			UserId:    cart.UserId,
			CartItems: cart.CartItems,
		}

		result, err := cartCollection.InsertOne(ctx, newCart)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.CartResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
