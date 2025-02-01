package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/okaaryanata/user/internal/domain"
)

func (c *Controller) CreateUser(ctx *gin.Context) {
	var (
		args domain.UserRequest
		err  error
	)

	defer func() {
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
	}()

	if err = ctx.ShouldBind(&args); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "errors": "Name is required"})
		return
	}

	user, err := c.userSvc.CreateUser(ctx, &args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": true, "user": user})
}

func (c *Controller) GetUsers(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
	}()

	pageNum, err := strconv.Atoi(ctx.DefaultQuery("page_num", "1"))
	if err != nil || pageNum < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "error": "Invalid page_num"})
		return
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "error": "Invalid page_size"})
		return
	}

	users, err := c.userSvc.GetUsers(ctx, &domain.GetUserRequest{
		Page: pageNum,
		Size: pageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"result": false, "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": true, "users": users})
}

func (c *Controller) GetUserByID(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
	}()

	id := ctx.Param("userID")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "errors": err.Error()})
		return
	}

	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "errors": "userID should be bigger than 0"})
		return
	}

	user, err := c.userSvc.GetUserByID(ctx, uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"result": false, "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": true, "user": user})
}
