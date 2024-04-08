package http

import (
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// user, err := domain.NewAuthenticatedUser(c)
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, "unauthorized a")
		// 	return
		// }

		// c.AbortWithStatusJSON(http.StatusUnauthorized, "unauthorized b")
		c.Next()
	}
}
