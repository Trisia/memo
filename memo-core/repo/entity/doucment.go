package entity

import (
	"encoding/json"
	"time"
)

// Document 文章
type Document struct {
	ID        uint      `gorm:"autoIncrement,type:bigint" json:"id"`
	CreatedAt time.Time `gorm:"type:datetime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updatedAt"`
	Creator   uint      `gorm:"type:bigint" json:"creator"`
	Title     string    `gorm:"type:varchar(256)" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Brief     string    `gorm:"type:varchar(516)" json:"brief"` // 文档摘要
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
