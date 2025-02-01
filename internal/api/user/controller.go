package user

import "github.com/gin-gonic/gin"

type (
	Controller struct {
	}
)

func NewUserController() *Controller {
	return &Controller{}
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	router.POST(``, c.CreateUser)
}
