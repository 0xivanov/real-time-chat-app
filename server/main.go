package main

import (
	"log"
	"server/controller"
	"server/db"
	"server/model"
	"server/repo"
	"server/router"
	"server/service"
)

func main() {
	dbConnection, err := db.NewDatabse("postgres", "postgresql://root:password@localhost:5433/go-chat")
	if err != nil {
		log.Fatalf("could not init database connection")
		panic(err)
	}
	userRepo := repo.NewUserRepository(dbConnection.GetDb())
	roomRepo := repo.NewRoomRepository(dbConnection.GetDb())
	userService := service.NewService(*userRepo)
	userController := controller.NewUserController(userService)
	rooms, err := roomRepo.GetRooms(dbConnection.GetDb().Statement.Context)
	if err != nil {
		log.Fatalf("could not get rooms: %v", err)
		panic(err)
	}

	hub := model.NewHub(rooms)
	wsController := controller.NewWsController(hub, roomRepo)
	go hub.Run()
	router.InitRouter(userController, wsController)
	router.Start("0.0.0.0:8000")
}
