package http

import (
	appProduct "loyalid-test/application/product"
	appUser "loyalid-test/application/user"

	"github.com/gin-gonic/gin"
)

func InitRoutes(
	userService *appUser.Service,
	productService *appProduct.Service,
) *gin.Engine {
	router := gin.Default()

	userCtrl := NewUserHandler(userService)
	router.GET("/user/current", Authenticate(), userCtrl.CurrentUser)

	productCtrl := NewProductHandler(productService)
	router.POST("/product", Authenticate(), productCtrl.CreateProduct)
	router.GET("/product", Authenticate(), productCtrl.ListProduct)

	return router
}
