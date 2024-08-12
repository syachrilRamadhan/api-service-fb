package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/products/:id", productsHandler)
	router.GET("/q", queryHandler)
	router.POST("/products", postProductsHandler)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to my RESTful API!",
		"version": "1.0.0",
		"author":  "Syachril Ramadhan",
	})
}

func productsHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Product details for ID " + id,
	})
}

func queryHandler(c *gin.Context) {
	product := c.Query("product")
	c.JSON(http.StatusOK, gin.H{
		"message": "Product details for product " + product,
	})
}

type Product struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required,number"`
}

func postProductsHandler(c *gin.Context) {
	var inputProduct Product

	err := c.ShouldBindJSON(&inputProduct)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": inputProduct.Title,
		"price": inputProduct.Price,
	})
}
