package controller

import (
	"log"
	"net/http"
	"server/model"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(s *service.UserService) *UserController {
	return &UserController{
		service: s,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user *model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.service.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	userID := ctx.Query("id")
	username := ctx.Query("username")
	log.Print(userID)
	log.Print(username)
	var err error
	var user *model.User
	if userID != "" {
		userIDint, err := strconv.Atoi(userID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err = c.service.GetUserById(ctx, userIDint)
	} else if username != "" {
		user, err = c.service.GetUserByUsername(ctx, username)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Login(ctx *gin.Context) {
	var user *model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loginResp, err := c.service.Login(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.SetCookie("jwt", loginResp.AccessToken, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, loginResp)
}

func (c *UserController) Logout(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

func (c *UserController) WebSocker(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
