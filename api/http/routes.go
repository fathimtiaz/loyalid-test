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
	router.GET("/user/current", Authenticate(), SetAuthdUserCtx(), userCtrl.CurrentUser)

	productCtrl := NewProductHandler(productService)
	router.POST("/product", Authenticate(), SetAuthdUserCtx(), productCtrl.CreateProduct)
	router.GET("/product", Authenticate(), SetAuthdUserCtx(), productCtrl.ListProduct)

	return router
}
