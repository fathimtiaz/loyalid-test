package http

import (
	"log"
	appProduct "loyalid-test/application/product"
	"loyalid-test/domain"
	"loyalid-test/repository"
	"net/http"
	"net/url"
	"strconv"

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

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var err error
	var product domain.Product

	if err = c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	if err = h.ProductService.CreateProduct(c.Request.Context(), &product); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) ListProduct(c *gin.Context) {
	var err error
	var query url.Values = c.Request.URL.Query()
	var result []domain.Product

	if result, err = h.ProductService.GetProducts(c.Request.Context(), getProductsQueryToFilter(query)); err != nil {
		if err == repository.ErrNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func getProductsQueryToFilter(query url.Values) (result repository.ProductFilter) {
	var err error
	var limit, page int

	if limit, err = strconv.Atoi(query.Get("limit")); err != nil {
		log.Print("error parsing query limit", err.Error())
	}

	if page, err = strconv.Atoi(query.Get("page")); err != nil {
		log.Print("error parsing query page", err.Error())
	}

	result.Pagination = repository.NewPagination(page, limit)

	return
}
