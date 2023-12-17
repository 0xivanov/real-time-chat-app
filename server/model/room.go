package model

type Room struct {
	ID      string             `json:"id" gorm:"primaryKey"`
	Name    string             `json:"name" gorm:"unique"`
	Clients map[string]*Client `json:"clients,omitempty" gorm:"-"`
}
