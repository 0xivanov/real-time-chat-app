package controller

import (
	"net/http"
	"server/model"
	"server/repo"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WsController struct {
	hub      *model.Hub
	roomRepo *repo.RoomRepo
}

func NewWsController(hub *model.Hub, roomRepo *repo.RoomRepo) *WsController {
	return &WsController{
		roomRepo: roomRepo,
		hub:      hub,
	}
}

func (c *WsController) CreateRoom(ctx *gin.Context) {
	var room *model.Room
	if err := ctx.ShouldBindJSON(&room); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.hub.Rooms[room.ID] = &model.Room{
		ID:      room.ID,
		Name:    room.Name,
		Clients: make(map[string]*model.Client),
	}
	err := c.roomRepo.CreateRoom(ctx, c.hub.Rooms[room.ID])
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, room)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// origin:= r.Header.Get("Origin")
		return true
	},
}

func (c *WsController) JoinRoom(ctx *gin.Context) {
	println("JOINROOM")
	roomID := ctx.Param("roomId")
	clientID := ctx.Query("userId")
	username := ctx.Query("username")

	if _, ok := c.hub.Rooms[roomID]; !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "room with this id does not exist"})
		return
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := &model.Client{
		Conn:     conn,
		Message:  make(chan *model.Message),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	message := &model.Message{
		Content:  "user has joined the room",
		Username: username,
		RoomID:   roomID,
	}

	// register a new client to the register channel
	c.hub.Register <- client
	c.hub.Broadcast <- message
	go client.WriteMessage()
	client.ReadMessage(c.hub)
}

func (c *WsController) GetRooms(ctx *gin.Context) {
	println("hello")
	ctx.JSON(http.StatusOK, c.hub.Rooms)
	// rooms, err := c.roomRepo.GetRooms(ctx)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, rooms)
}
