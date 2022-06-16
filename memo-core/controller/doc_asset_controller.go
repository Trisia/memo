package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"memo-core/objstore"
	"net/http"
	"os"
	"strconv"
)

func NewDocAssetController(r gin.IRouter) *DocAssetController {
	res := &DocAssetController{}
	base := r.Group("doc")
	// 上传文档资源
	base.POST(":id/asset", res.putAsset)
	// 下载文档资源
	base.GET(":id/asset/:filename", res.getAsset)
	return res
}

type DocAssetController struct {
}

/**
@api {PUT} /api/doc/:id/asset 上传资源
@apiDescription 上传于文档的资源，如：图片、文档等
@apiName DocPutAsset
@apiGroup Doc

@apiHeader Content-Type multipart/form-data

@apiParam {File} file 上传文件。

@apiSuccess {String} Body 文件存储名称

@apiSuccessExample {HTTP} 成功响应
HTTP/1.1 200

f7636c54c6fa48b59afa1a96006e782cf7636c54c6fa48b59afa1a96006e782c.png

@apiErrorExample {HTTP} 失败
HTTP/1.1 400

无法解析参数
*/
func (c *DocAssetController) putAsset(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	if id <= 0 {
		ErrIllegal(ctx, "文档ID错误")
		return
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		ErrIllegal(ctx, "文件上传请求格式错误或文件为空")
		return
	}
	src, err := file.Open()
	if err != nil {
		ErrSys(ctx, err)
		return
	}
	defer src.Close()
	hashName, _, err := objstore.Repo.Put(idStr, file.Filename, src)
	if err != nil {
		ErrSys(ctx, err)
		return
	}
	ctx.String(200, hashName)
}

/**
@api {GET} /api/doc/:id/asset/:filename 下载资源
@apiDescription 下载文档相关的资源，如：图片、文档等
@apiName DocGetAsset
@apiGroup Doc

@apiHeader Content-Type application/octet-stream

@apiSuccess {Binary} Body 资源文件内容

@apiErrorExample {HTTP} 失败
HTTP/1.1 404
*/
func (c *DocAssetController) getAsset(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	if id <= 0 {
		ErrNotFount(ctx)
		return
	}
	filename := ctx.Param("filename")
	if filename == "" {
		ErrNotFount(ctx)
		return
	}
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	ctx.Header("Content-Type", "application/octet-stream")
	_, err := objstore.Repo.Get(idStr, filename, ctx.Writer)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			ErrNotFount(ctx)
			return
		} else {
			ErrSys(ctx, err)
			return
		}
	}
}
