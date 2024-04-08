package http

import (
	appProduct "loyalid-test/application/product"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService *appProduct.Service
}

func NewProductHandler(productService *appProduct.Service) *ProductHandler {
	return &ProductHandler{
		ProductService: productService,
	}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {

}

func (h *ProductHandler) ListProduct(ctx *gin.Context) {

}
