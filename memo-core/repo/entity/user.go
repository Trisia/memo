package entity

import (
	"encoding/json"
	"time"
)

// User 用户
type User struct {
	ID        uint      `gorm:"autoIncrement,type:bigint" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `gorm:"type:varchar(64)" json:"username"`
	Email     string    `gorm:"type:varchar(256)" json:"email"`
	Password  string    `gorm:"type:varchar(256)" json:"-"`
	Salt      string    `gorm:"type:varchar(256)" json:"-"`
	Typ       int       `gorm:"type:tinyint" json:"typ"` // 用户类型 0 - 普通用户； 1 - 管理员
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
