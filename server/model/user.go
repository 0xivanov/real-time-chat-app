package model

type User struct {
	ID       int    `json:"id,omitempty" gorm:"primaryKey"`
	Username string `json:"username,omitempty" gorm:"unique"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

type LoginUserResp struct {
	AccessToken string `json:"access_token"`
	ID          int    `json:"id,omitempty"`
	Username    string `json:"username,omitempty"`
}
