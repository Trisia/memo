package repo

import (
	"gorm.io/gorm"
	"memo-core/repo/entity"
)

func NewDocTagAggRepo() *DocTagAggRepo {
	return &DocTagAggRepo{}
}

type DocTagAggRepo struct {
}

// Save 保存或更新文档以及tag属性
func (r *DocTagAggRepo) Save(record *entity.Document, tags []uint) error {
	if record == nil {
		return nil
	}
	return DB.Transaction(func(tx *gorm.DB) error {
		var err error
		if record.ID <= 0 {
			// 创建文档
			err = tx.Save(&record).Error
		} else {
			// 更新文档
			err = tx.Model(&record).Select("title", "content").Updates(record).Error
		}
		if err != nil {
			return err
		}
		// 删除旧的文档ID关系
		docId := record.ID
		err = tx.Where("doc_id = ?", docId).Delete(&entity.DocTag{}).Error
		if err != nil {
			return err
		}
		// 创建新的文档ID与文档的关系
		if len(tags) > 0 {
			newTags := make([]entity.DocTag, len(tags))
			for i, tagId := range tags {
				newTags[i] = entity.DocTag{DocId: docId, TagId: tagId}
			}
			err = tx.Create(&newTags).Error
		}
		return err
	})
}
