package dto

type TagCountItem struct {
	ID    uint   `json:"id"`    // 标签记录ID
	Name  string `json:"name"`  // 标签名称
	Count int    `json:"count"` // 该标签下文档数量
}
