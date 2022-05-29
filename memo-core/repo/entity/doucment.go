package entity

import (
	"encoding/json"
	"time"
)

// Document 文章
type Document struct {
	ID        uint      `gorm:"autoIncrement,type:bigint" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Creator   uint      `gorm:"type:bigint" json:"creator"`
	Title     string    `gorm:"type:varchar(256)" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
}

func (c *Document) MarshalJSON() ([]byte, error) {
	type Alias Document
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
