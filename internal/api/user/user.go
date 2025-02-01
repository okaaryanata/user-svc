package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"result": true, "users": "users"})
}
