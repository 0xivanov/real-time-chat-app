package repo

import (
	"context"
	"fmt"
	"server/model"

	"gorm.io/gorm"
)

type RoomRepo struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepo {
	return &RoomRepo{
		db: db,
	}
}

func (repo *RoomRepo) CreateRoom(ctx context.Context, room *model.Room) error {
	if err := repo.db.Create(room).Error; err != nil {
		return fmt.Errorf("could not create rooms: %v", err)
	}
	return nil
}

func (repo *RoomRepo) GetRooms(ctx context.Context) (map[string]*model.Room, error) {
	var rooms []model.Room
	if err := repo.db.WithContext(ctx).Find(&rooms).Error; err != nil {
		return nil, fmt.Errorf("could not find rooms: %v", err)
	}
	var roomsMap = make(map[string]*model.Room)
	for _, v := range rooms {
		v.Clients = make(map[string]*model.Client)
		roomsMap[v.ID] = &v
	}
	return roomsMap, nil
}
