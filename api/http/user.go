package http

import (
	appUser "loyalid-test/application/user"
	"loyalid-test/domain"
	"loyalid-test/repository"
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

	if user, err = h.UserService.CurrentUser(c.Request.Context()); err != nil {
		if err == repository.ErrNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}
