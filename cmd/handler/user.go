package handler

import (
	"net/http"

	"github.com/job25721/go-jwt/pkg/model"
	"github.com/job25721/go-jwt/pkg/user"

	"github.com/gin-gonic/gin"
)

type handler struct {
	userService user.IUserService
}

func NewUserHandler(usersrv user.IUserService) handler {
	return handler{
		userService: usersrv,
	}
}

func (h *handler) RegisterHandler(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bad Request")
		return
	}

	if req.Password == "" || req.Username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bad Request")
		return
	}

	resp, status, err := h.userService.Register(req)

	if err != nil {
		c.AbortWithStatusJSON(status, err.Error())
		return
	}

	c.JSON(status, resp)
}

func (h *handler) Login(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bad Request")
		return
	}

	if req.Password == "" || req.Username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bad Request")
		return
	}

	token, status, err := h.userService.Login(req)

	if err != nil {
		c.AbortWithStatusJSON(status, err.Error())
		return
	}

	c.JSON(status, gin.H{"token": token})
}
