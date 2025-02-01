package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetUsers(ctx *gin.Context) {
	users, err := c.userSvc.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": true, "users": users})
}
