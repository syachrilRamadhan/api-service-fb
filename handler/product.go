package handler

import (
	"api-service-fb/product"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to my RESTful API!",
		"version": "1.0.0",
		"author":  "Syachril Ramadhan",
	})
}

func ProductsHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Product details for ID " + id,
	})
}

func QueryHandler(c *gin.Context) {
	product := c.Query("product")
	c.JSON(http.StatusOK, gin.H{
		"message": "Product details for product " + product,
	})
}

func PostProductsHandler(c *gin.Context) {
	var inputProduct product.Product

	err := c.ShouldBindJSON(&inputProduct)
	if err != nil {
		// cek apakah error karena validasi
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := []string{}
			for _, e := range validationErrors {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}

		if unmarshalError, ok := err.(*json.UnmarshalTypeError); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Invalid type for field %s, expected %s", unmarshalError.Field, unmarshalError.Type),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "An unexpected error occurred",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": inputProduct.Title,
		"price": inputProduct.Price,
	})
}
