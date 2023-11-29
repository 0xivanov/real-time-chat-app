package main

import (
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.GET("/ws", func(c *gin.Context) {
		if err := m.HandleRequest(c.Writer, c.Request); err != nil {
			fmt.Println(err)
			return
		}
	})

	m.HandleConnect(func(s *melody.Session) {
		fmt.Println("New WebSocket connection established")
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Run(":8000")
}
