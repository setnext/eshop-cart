package controllers

import (
	"context"
	"eshop-cart-api/configs"
	"eshop-cart-api/models"
	"eshop-cart-api/responses"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection = configs.GetCollection(configs.DB, "carts")
var validate = validator.New()

func CreateOrUpdateCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var cart models.Cart
		// var cart1 models.Cart
		var existingCart models.Cart
		defer cancel()

		fmt.Println("Received the call")
		//validate the request body

		// var rr = c.BindJSON(&cart1)
		// fmt.Println(rr)

		fmt.Println(cart)
		if err := c.BindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, responses.CartResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&cart); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.CartResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		var userId = cart.UserId
		err := cartCollection.FindOne(ctx, bson.M{"userId": userId}).Decode(&existingCart)
		if err != nil {
			fmt.Println("Record Not Found")

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
			return
		}

		fmt.Println("Cart Found, So Updating the Cart")

		update := bson.M{"cartItems": cart.CartItems}
		result, err := cartCollection.UpdateOne(ctx, bson.M{"userId": cart.UserId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		fmt.Println("Update Completed")
		fmt.Println(cart.UserId)

		//get updated user details
		var updatedCart models.Cart
		if result.MatchedCount == 1 {
			err := cartCollection.FindOne(ctx, bson.M{"userId": cart.UserId}).Decode(&updatedCart)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedCart}})

	}
}

func GetCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("uid")
		var user models.Cart
		defer cancel()

		err := cartCollection.FindOne(ctx, bson.M{"userId": userId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "No Cart Found for the User"}})
			return
		}

		c.JSON(http.StatusOK, responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

func DeleteCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("uid")
		defer cancel()

		// objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := cartCollection.DeleteOne(ctx, bson.M{"userId": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.CartResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Cart with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Cart successfully deleted!"}},
		)
	}
}

func GetAllCarts() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var carts []models.Cart
		defer cancel()
		fmt.Println("Retrieveing all carts")
		results, err := cartCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUserCart models.Cart
			if err = results.Decode(&singleUserCart); err != nil {
				c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			carts = append(carts, singleUserCart)
		}

		c.JSON(http.StatusOK,
			responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": carts}},
		)
	}
}
