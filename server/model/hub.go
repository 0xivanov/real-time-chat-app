package model

import "fmt"

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub(rooms map[string]*Room) *Hub {
	return &Hub{
		Rooms:      rooms,
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if _, ok := h.Rooms[client.RoomID]; ok {
				room := h.Rooms[client.RoomID]
				if _, ok := room.Clients[client.ID]; !ok {
					room.Clients[client.ID] = client
				}
			}
		case client := <-h.Unregister:
			if _, ok := h.Rooms[client.RoomID]; ok {
				println("USER LEFT2")
				room := h.Rooms[client.RoomID]
				if _, ok := room.Clients[client.ID]; ok {
					println("USER LEFT1")
					if len(h.Rooms[client.RoomID].Clients) != 0 {
						println("USER LEFT3")
						// broadcast that the client has left the room
						// h.Broadcast <- &Message{
						// 	Content:  "user has left the chat",
						// 	Username: client.Username,
						// 	RoomID:   client.RoomID,
						// }
					}
					println(room.Clients)
					println(client.ID)
					delete(room.Clients, client.ID)
					close(client.Message)
					println("USER LEFT")
				}
			}
		case message := <-h.Broadcast:
			fmt.Println(message)
			fmt.Println(message.Content)
			if _, ok := h.Rooms[message.RoomID]; ok {
				for _, client := range h.Rooms[message.RoomID].Clients {
					client.Message <- message
				}
			}
		}
	}
}
