package dto

// DocSaveParam 保存文档所需的参数
type DocSaveParam struct {
	ID           uint     `json:"id"`           // 文档记录ID，存在时表示更新否则为删除
	Creator      uint     `json:"creator"`      // 创建者用户ID
	Title        string   `json:"title"`        // 文档标题
	Content      string   `json:"content"`      // 文档内容
	Tags         []uint   `json:"tags"`         // 文档关联的标签
	DeleteImages []string `json:"deleteImages"` // 待删除文档
}
