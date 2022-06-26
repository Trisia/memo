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
	Password  string    `gorm:"type:varchar(256)" json:"password"`
	Salt      string    `gorm:"type:varchar(256)" json:"-"`
	Typ       int       `gorm:"type:tinyint" json:"typ"` // 用户类型 0 - 普通用户； 1 - 管理员; 2 - 应用
	Avatar    []byte    `gorm:"type:blob" json:"avatar"` // 头像
}

type user struct {
	User
	CreatedAt DateTime `json:"createdAt"`
	UpdatedAt DateTime `json:"updatedAt"`
}

func (c *User) MarshalJSON() ([]byte, error) {
	e := user{
		User:      *c,
		CreatedAt: DateTime(c.CreatedAt),
		UpdatedAt: DateTime(c.UpdatedAt),
	}
	e.Password = ""
	return json.Marshal(e)
}
