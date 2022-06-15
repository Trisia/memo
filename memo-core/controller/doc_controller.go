package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"memo-core/controller/dto"
	"memo-core/repo"
	"memo-core/repo/entity"
	"strconv"
)

// DocController 文档接口
type DocController struct {
}

func NewDocController(r gin.IRouter) *DocController {
	res := &DocController{}
	base := r.Group("doc")
	// 创建或更新文档内容
	base.POST("", res.save)
	// 删除文档
	base.DELETE("", res.delete)
	// 获取文档内容
	base.GET(":id", res.get)
	// 查询文档
	base.GET("search", res.search)

	return res
}

/**
@api {POST} /api/doc 保存文档
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

@apiSuccessExample {HTTP} 成功响应
HTTP/1.1 200

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

/**
@api {DELETE} /api/doc 删除文档
@apiDescription 通过用户ID删除文档。
@apiName DocDelete
@apiGroup Doc

@apiParam {Integer} [id] 文档ID，ID存在时为更新

@apiParamExample {HTTP} 删除
DELETE /doc?id=1

@apiSuccessExample {HTTP} 成功响应
HTTP/1.1 200

@apiErrorExample {http} 失败
HTTP/1.1 500

内部错误
*/
func (c *DocController) delete(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, _ := strconv.Atoi(idStr)
	if id <= 0 {
		ErrIllegal(ctx, "文档ID错误")
		return
	}
	userId := ctx.GetUint("userId")

	err := repo.DB.Where("creator = ? AND id = ?", userId, id).Delete(&entity.Document{}).Error
	if err != nil {
		ErrSys(ctx, err)
		return
	}
}

/**
@api {GET} /api/doc/:id 获取文档
@apiDescription 通过文档ID获取文档
@apiName DocGet
@apiGroup Doc

@apiParam {Integer} id 文档ID
@apiParamExample {HTTP} 删除
GET /doc/1

@apiSuccess {Integer} id 文档ID
@apiSuccess {String} title 文档标题
@apiSuccess {String} content 文档内容
@apiSuccess {String} createdAt 创建时间，YYYY-MM-DD HH:mm:ss
@apiSuccess {String} updatedAt 更新时间，YYYY-MM-DD HH:mm:ss
@apiSuccess {Integer} creator 文档所属用户ID

@apiSuccessExample {HTTP} 成功响应
HTTP/1.1 200

{
    "id": 1,
    "creator": 7,
    "title": "标题2",
    "content": "# 标题2",
    "createdAt": "2022-06-09 22:23:49",
    "updatedAt": "2022-06-11 19:12:37"
}

@apiErrorExample {http} 失败
HTTP/1.1 400

文档ID错误
*/
func (c *DocController) get(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	if id <= 0 {
		ErrIllegal(ctx, "文档ID错误")
		return
	}
	var res entity.Document
	if err := repo.DB.First(&res, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ErrIllegal(ctx, "文档ID错误")
			return
		} else {
			ErrSys(ctx, err)
			return
		}
	}
	ctx.JSON(200, &res)
}

/**
@api {GET} /api/doc/search 查询文档
@apiDescription 批量查询文档信息
@apiName DocSearch
@apiGroup Doc

@apiParam {Integer} [limit=30] 本次查询返回最大记录数记录。
@apiParam {Integer} [offset=0] 偏移量
@apiParam {String} [keyword] 标题关键字
@apiParamExample {HTTP} 删除
GET /doc/search?size=5&after=12&keyword=标题

@apiSuccess {Doc[]} Body 文档数据
@apiSuccess {Integer} Doc.id 文档ID
@apiSuccess {String} Doc.title 文档标题
@apiSuccess {String} Doc.content 文档内容
@apiSuccess {String} Doc.createdAt 创建时间，YYYY-MM-DD HH:mm:ss
@apiSuccess {String} Doc.updatedAt 更新时间，YYYY-MM-DD HH:mm:ss
@apiSuccess {Integer} Doc.creator 文档所属用户ID

@apiSuccessExample {HTTP} 成功响应
HTTP/1.1 200

[
    {
        "id": 2,
        "creator": 7,
        "title": "文档1",
        "content": "# 标题",
        "createdAt": "2022-06-11 20:01:45",
        "updatedAt": "2022-06-11 20:01:45"
    },
    {
        "id": 1,
        "creator": 7,
        "title": "标题2",
        "content": "# 标题2",
        "createdAt": "2022-06-09 22:23:49",
        "updatedAt": "2022-06-11 19:12:37"
    }
]
@apiErrorExample {HTTP} 失败
HTTP/1.1 400

无法解析参数
*/
func (c *DocController) search(ctx *gin.Context) {
	var param dto.DocSearch
	if err := ctx.ShouldBindQuery(&param); err != nil {
		ErrIllegal(ctx, "无法解析参数")
		return
	}
	if param.Limit <= 0 {
		param.Limit = 30
	}
	if param.Offset <= 0 {
		param.Offset = 0
	}
	userId := ctx.GetUint("userId")

	tx := repo.DB.
		Limit(param.Limit).
		Offset(param.Offset).
		Order("created_at desc").
		Where("creator = ?", userId)
	if param.Keyword != "" {
		tx = tx.Where("title LIKE ?", fmt.Sprintf("%%%s%%", param.Keyword))
	}
	var res []*entity.Document
	if err := tx.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(0, []*entity.Document{})
			return
		} else {
			ErrSys(ctx, err)
			return
		}
	}
	ctx.JSON(200, res)
}
