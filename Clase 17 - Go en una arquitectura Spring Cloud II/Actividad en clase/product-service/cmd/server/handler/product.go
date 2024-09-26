package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tomas.pereyra/product-service/internal/domain"
	"gitlab.com/tomas.pereyra/product-service/internal/product"
)

type ProductHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *ProductHandler {
	return &ProductHandler{service: service}

}

func (productHandler *ProductHandler) FindById() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, productHandler.service.FindProductById(c))
	}

}
func (productHandler *ProductHandler) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product domain.Product
		c.ShouldBind(&product)
		c.JSON(200, productHandler.service.SaveProduct(product))
	}

}
