package db

import (
	"server/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabse(dbDriver string, dbConnectionString string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Room{})

	return &Database{db: db}, nil
}

func (d *Database) GetDb() *gorm.DB {
	return d.db
}
