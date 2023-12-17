package router

import (
	"server/controller"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(uc *controller.UserController, wc *controller.WsController) {
	r = gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", uc.CreateUser)
	r.POST("/login", uc.Login)
	r.GET("/logout", uc.Logout)
	r.GET("/user", uc.GetUser)
	r.POST("/ws/createRoom", wc.CreateRoom)
	r.GET("ws/joinRoom/:roomId", wc.JoinRoom)
	r.GET("ws/getRooms", wc.GetRooms)
}

func Start(addr string) error {
	return r.Run()
}
