package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/job25721/go-jwt/pkg/jwt"
)

const JwtCliams = "user"

func NewJwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
		cliams, err := jwt.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		c.Set(JwtCliams, cliams)
	}
}
