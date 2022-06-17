package controller

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"memo-core/controller/midd"
	"memo-core/repo"
	"memo-core/repo/entity"
)

func NewTagController(r gin.IRouter) *TagController {
	res := &TagController{}
	base := r.Group("tag", midd.Entity)
	base.POST("", res.create)
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
