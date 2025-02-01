package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/okaaryanata/user/internal/domain"
)

func (c *Controller) CreateUser(ctx *gin.Context) {
	var args domain.UserRequest
	if err := ctx.ShouldBind(&args); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	user, err := c.userSvc.CreateUser(ctx, &args)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": true, "user": user})
}

func (c *Controller) GetUsers(ctx *gin.Context) {
	users, err := c.userSvc.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": true, "users": users})
}

func (c *Controller) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("userID")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	user, err := c.userSvc.GetUserByID(ctx, uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": true, "user": user})
}
