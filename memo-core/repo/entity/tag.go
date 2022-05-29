package entity

import (
	"encoding/json"
	"time"
)

// Tag 文章的标签
type Tag struct {
	ID        uint      `gorm:"autoIncrement,type:bigint" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Creator   uint      `gorm:"type:bigint" json:"creator"`
	Value     string    `gorm:"type:varchar(512)" json:"value"`
}

func (c *Tag) MarshalJSON() ([]byte, error) {
	type Alias Tag
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
