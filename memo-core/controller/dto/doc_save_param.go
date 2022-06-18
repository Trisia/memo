package dto

import "memo-core/repo/entity"

// DocSaveParam 保存文档所需的参数
type DocSaveParam struct {
	entity.Document
	Tags      []uint   `json:"tags"`      // 文档关联的标签序列
	DeleteImg []string `json:"deleteImg"` // 待删除的图片名称序列
}
