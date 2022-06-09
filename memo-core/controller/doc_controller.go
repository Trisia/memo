package controller

import (
	"github.com/gin-gonic/gin"
	"memo-core/repo"
	"memo-core/repo/entity"
)

// DocController 文档接口
type DocController struct {
}

func NewDocController(r gin.IRouter) *DocController {
	res := &DocController{}
	base := r.Group("doc")
	base.POST("save", res.save)
	return res
}

/**
@api {POST} /doc/save 保存文档
@apiDescription 保存文档，如果ID为空则为创建，否则为更新。
@apiName DocSave
@apiGroup Doc

@apiParam {Integer} [id] 文档ID，ID存在时为更新
@apiParam {String} title 文档标题
@apiParam {String} [content] 文档内容

@apiParamExample {json} 创建
{
	"title": "标题",
	"content": "#标题"
}
@apiParamExample {json} 更新
{
	"id": 1,
	"title": "标题2",
	"content": "#标题2"
}
@apiSuccess {Integer} Body 文档ID

@apiSuccessExample {Integer} 成功响应
1
@apiErrorExample {http} 失败
HTTP/1.1 400

文档内容(content)为空
*/
func (c *DocController) save(ctx *gin.Context) {
	var param entity.Document
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrIllegal(ctx, "无法解析参数")
		return
	}
	param.Creator = ctx.GetUint("userId")
	if param.Title == "" {
		ErrIllegal(ctx, "文档标题(title)为空")
		return
	}

	var err error
	// 创建
	if param.ID <= 0 {
		err = repo.DB.Save(&param).Error
	} else {
		err = repo.DB.Model(&param).Select("title", "content").Updates(param).Error
	}
	if err != nil {
		ErrSys(ctx, err)
		return
	}
	ctx.JSON(200, param.ID)

}
