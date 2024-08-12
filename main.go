package main

import (
	"api-service-fb/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handler.RootHandler)
	router.GET("/products/:id", handler.ProductsHandler)
	router.GET("/q", handler.QueryHandler)
	router.POST("/products", handler.PostProductsHandler)

	router.Run()
}
