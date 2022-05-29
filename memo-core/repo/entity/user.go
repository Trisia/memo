package entity

import (
	"encoding/json"
	"time"
)

// User 用户
type User struct {
	ID        uint `gorm:"autoIncrement,type:bigint" `
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `gorm:"type:varchar(64)"`
	Email     string `gorm:"type:varchar(256)"`
	Password  string `gorm:"type:varchar(256)"`
	Salt      string `gorm:"type:varchar(256)"`
}

func (c *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		*Alias
		CreatedAt DateTime `json:"createdAt"`
		UpdatedAt DateTime `json:"updatedAt"`
	}{
		(*Alias)(c),
		DateTime(c.CreatedAt),
		DateTime(c.UpdatedAt),
	})
}
