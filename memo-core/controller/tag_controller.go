package controller

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"memo-core/controller/dto"
	"memo-core/controller/midd"
	"memo-core/repo"
	"memo-core/repo/entity"
)

func NewTagController(r gin.IRouter) *TagController {
	res := &TagController{}
	base := r.Group("tag", midd.Entity)
	base.POST("", res.create)
	// 聚合列表用户标签
	base.GET("list", res.list)
	return res
}

type TagController struct {
}

/**
@api {POST} /api/tag 创建标签
@apiDescription 创建新的标签，同一用户标签不可重复。
@apiName TagCreate
@apiGroup Tag

@apiParam {String} body 标签，不能重复

@apiParamExample {HTTP} 创建
POST /api/tag

笔记

@apiSuccess {Integer} Body 标签ID
@apiSuccessExample {HTTP} 成功响应
HTTP/1.1 200

1

@apiErrorExample {http} 失败
HTTP/1.1 400

标签已存在
*/
func (c *TagController) create(ctx *gin.Context) {
	b, _ := ioutil.ReadAll(ctx.Request.Body)
	if len(b) == 0 {
		ErrIllegal(ctx, "标签不能为空")
		return
	}

	value := string(b)
	userId := ctx.GetUint("userId")
	var cnt int64
	err := repo.DB.Model(&entity.Tag{}).Where("value = ? AND creator = ?", value, userId).Count(&cnt).Error
	if err != nil {
		ErrSys(ctx, err)
		return
	}
	if cnt > 0 {
		ErrIllegal(ctx, "标签已存在")
		return
	}

	record := entity.Tag{Creator: userId, Value: value}
	err = repo.DB.Create(&record).Error
	if err != nil {
		ErrSys(ctx, err)
		return
	}
	ctx.JSON(200, record.ID)
}

/**
@api {GET} /api/tag/list 标签列表
@apiDescription 获取用户标签列表
@apiName TagList
@apiGroup Tag

@apiParamExample {HTTP} 列表
GET /api/tag/list

@apiSuccess {Item[]} body 标签数组
@apiSuccess {Object} Item 标签对象
@apiSuccess {Integer} Item.id 标签记录ID
@apiSuccess {String} Item.name 标签名称
@apiSuccess {Integer} Item.count 文档数量
@apiSuccessExample {json} 成功响应
HTTP/1.1 200

[
    {
        "id": 1,
        "name": "笔记",
        "count": 3
    },
    {
        "id": 2,
        "name": "Docker",
        "count": 1
    }
]

@apiErrorExample {http} 失败
HTTP/1.1 500

内部错误
*/
func (c *TagController) list(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	_ = userId
	var tagArr []entity.Tag
	err := repo.DB.Where("creator = ?", userId).Find(&tagArr).Error
	if err != nil {
		ErrSys(ctx, err)
		return
	}
	if len(tagArr) == 0 {
		ctx.JSON(200, []dto.TagCountItem{})
		return
	}
	var m = map[uint]*dto.TagCountItem{}
	res := make([]dto.TagCountItem, len(tagArr))
	tagIdArr := make([]uint, len(tagArr))
	for i, item := range tagArr {
		tagIdArr[i] = item.ID
		res[i] = dto.TagCountItem{ID: item.ID, Name: item.Value, Count: 0}
		m[item.ID] = &res[i]
	}

	var contRes []dto.TagCountItem
	// SELECT tag_id as id , COUNT(doc_id) as count FROM doc_tags WHERE tag_id IN (1,2,3,4,5,6) GROUP BY tag_id;
	err = repo.DB.Model(&entity.DocTag{}).Select("tag_id as id, COUNT(doc_id) as count").Where("tag_id IN(?)", tagIdArr).Group("tag_id").Find(&contRes).Error
	if err != nil {
		ErrSys(ctx, err)
		return
	}
	for _, item := range contRes {
		if cntItem, ok := m[item.ID]; ok {
			cntItem.Count = item.Count
		}
	}
	ctx.JSON(200, res)
}
