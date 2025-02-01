package user

import (
	"github.com/gin-gonic/gin"

	"github.com/okaaryanata/user/internal/service"
)

type (
	Controller struct {
		userSvc *service.UserService
	}
)

func NewUserController(userSvc *service.UserService) *Controller {
	return &Controller{
		userSvc: userSvc,
	}
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	router.GET(``, c.GetUsers)
}
