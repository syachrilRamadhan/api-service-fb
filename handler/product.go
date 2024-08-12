package handler

import (
	"api-service-fb/product"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type produkHandler struct {
	produkService product.Service
}

func NewProdukHandler(produkService product.Service) *produkHandler {
	return &produkHandler{produkService}
}

func (h *produkHandler) GetProducts(c *gin.Context) {
	produk, err := h.produkService.GetProduk()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    produk,
	})
}

func (h *produkHandler) PostProductsHandler(c *gin.Context) {
	var productRequest product.ProdukRequest

	err := c.ShouldBindJSON(&productRequest)
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

	produk, err := h.produkService.CreateProduk(productRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  200,
		"message": "create product successfully",
		"data":    produk,
	})
}
