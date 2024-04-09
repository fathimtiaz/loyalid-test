package http

import (
	appUser "loyalid-test/application/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *appUser.Service
}

func NewUserHandler(userService *appUser.Service) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {

}

func (h *UserHandler) Authenticate(ctx *gin.Context) {

}

func (h *UserHandler) CurrentUser(ctx *gin.Context) {

}
