package http

import (
	appUser "loyalid-test/application/user"
	"loyalid-test/domain"
	"net/http"

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

func (h *UserHandler) CurrentUser(c *gin.Context) {
	var err error
	var user domain.User

	if user, err = h.UserService.CurrentUser(c); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
