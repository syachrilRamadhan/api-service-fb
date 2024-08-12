package main

import (
	"api-service-fb/handler"
	"api-service-fb/product"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/family-battery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&product.Produk{})

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	produkHandler := handler.NewProdukHandler(productService)

	router := gin.Default()

	router.GET("/", produkHandler.RootHandler)
	router.GET("/products/:id", produkHandler.ProductsHandler)
	router.GET("/q", produkHandler.QueryHandler)
	router.POST("/products", produkHandler.PostProductsHandler)

	router.Run()
}
