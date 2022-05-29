package entity

import (
	"encoding/json"
	"time"
)

// DocTag 文章的标签关系表
type DocTag struct {
	ID        uint      `gorm:"autoIncrement,type:bigint" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DocId     uint      `gorm:"type:bigint" json:"docId"`
	TagId     uint      `gorm:"type:bigint" json:"tagId"`
}

func (c *DocTag) MarshalJSON() ([]byte, error) {
	type Alias DocTag
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
