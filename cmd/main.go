package main

import (
	"net/http"

	"github.com/job25721/go-jwt/cmd/handler"
	"github.com/job25721/go-jwt/pkg/middleware"
	"github.com/job25721/go-jwt/pkg/store"
	"github.com/job25721/go-jwt/pkg/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello go authentication jwt example")
	})

	store := store.NewStore()
	userService := user.NewService(store)
	userHandler := handler.NewUserHandler(userService)

	r.POST("/register", userHandler.RegisterHandler)
	r.POST("/login", userHandler.Login)

	api := r.Group("/secret", middleware.NewJwtMiddleware())
	{
		api.GET("/", func(c *gin.Context) {
			cliams, _ := c.Get(middleware.JwtCliams)
			c.JSON(http.StatusOK, cliams)
		})
	}

	r.Run(":3000")
}
