package controller

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/tjfoc/gmsm/sm3"
	"io/ioutil"
	"memo-core/controller/midd"
	"memo-core/repo"
	"memo-core/repo/entity"
	"net/mail"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	KB = 1 << (10 * 1)
	MB = 1 << (10 * 2)
)

func NewUserController(r gin.IRouter) *UserController {
	res := &UserController{}
	base := r.Group("user")
	// 创建用户
	base.POST("register", res.register)
	base.DELETE("", midd.AdminOnly, res.delete)

	// 查询用户信息
	base.GET("info", midd.UserOnly, res.info)
	// 上传头像
	base.POST("avatar", midd.UserOnly, res.uploadAvatar)
	return res
}

// UserController 用户控制器
type UserController struct {
}

/**
@api {POST} /api/user/register 注册用户
@apiDescription 注册用户。
@apiName UserCreate
@apiGroup User

@apiParam {String} username 用户名
@apiParam {String} email 邮箱
@apiParam {String} password

@apiParamExample {json} 请求示例
{
	"username": "user1",
	"email": "example@email.com",
	"password": "123qwe"
}

@apiSuccess {Integer} id 记录ID
@apiSuccess {String} username 用户名
@apiSuccess {String} email 邮箱
@apiSuccess {String} createdAt 创建时间，YYYY-MM-DD HH:mm:ss
@apiSuccess {String} updatedAt 更新时间，YYYY-MM-DD HH:mm:ss

@apiSuccessExample {json} 成功响应
{
    "id": 6,
    "username": "user3",
    "email": "example@email211.com",
    "password": "",
    "typ": 0,
    "createdAt": "2022-06-08 22:24:57",
    "updatedAt": "2022-06-08 22:24:57"
}
*/
func (c *UserController) register(ctx *gin.Context) {
	var param entity.User
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrIllegal(ctx, "无法解析参数")
		return
	}
	if param.Username == "" {
		ErrIllegal(ctx, "用户名为空")
		return
	}
	if _, err := mail.ParseAddress(param.Email); err != nil {
		ErrIllegal(ctx, "邮箱非法")
		return
	}
	if len(param.Password) < 8 {
		ErrIllegal(ctx, "口令长度不足8位")
		return
	}

	exist, err := repo.UserSvc.Exist(param.Username, param.Email)
	if err != nil {
		ErrSys(ctx, err)
		return
	}
	if exist {
		ErrIllegal(ctx, "用户名或邮箱已经存在")
		return
	}
	salt := make([]byte, 16)
	_, _ = rand.Reader.Read(salt)
	hash := sm3.New()
	hash.Write([]byte(param.Password))
	hash.Write(salt)
	password := hash.Sum(nil)
	param.Salt = hex.EncodeToString(salt)
	param.Password = hex.EncodeToString(password)
	param.Typ = 0
	err = repo.DB.Create(&param).Error
	if err != nil {
		ErrSys(ctx, err)
		return
	}
	ctx.JSON(200, &param)
}

/**
@api {DELETE} /api/user 删除用户
@apiDescription 删除用户。
@apiName UserDelete
@apiGroup User

@apiHeader {String} token 管理员认证令牌


@apiParam {Integer} id 记录ID


@apiParamExample {HTTP} 请求示例
DELETE /user?id=1

@apiSuccessExample {HTTP} 成功响应
HTTP/1.1 200 OK
*/
func (c *UserController) delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	if id <= 0 {
		ErrIllegal(ctx, "记录ID错误")
		return
	}
	if err := repo.DB.Delete(&entity.User{}, id).Error; err != nil {
		ErrSys(ctx, err)
		return
	}
}

/**
@api {GET} /api/user/info 用户信息
@apiDescription 查询用户信息。
@apiName UserInfo
@apiGroup User

@apiHeader {String} token 认证令牌

@apiParamExample {HTTP} 请求示例
GET /user/info

@apiSuccess {Integer} id 记录ID
@apiSuccess {String} username 用户名
@apiSuccess {String} email 邮箱
@apiSuccess {String} createdAt 创建时间，YYYY-MM-DD HH:mm:ss
@apiSuccess {String} updatedAt 更新时间，YYYY-MM-DD HH:mm:ss
@apiSuccess {String} [avatar] 头像的Base64编码
@apiSuccess {Integer} typ 用户类型， 0 - 普通用户、1 - 管理员、2 - 应用

@apiSuccessExample {HTTP} 成功响应
HTTP/1.1 200 OK

{
    "id": 2,
    "createdAt": "2022-06-29T20:54:10.15+08:00",
    "updatedAt": "2022-06-29T20:54:10.15+08:00",
    "username": "test",
    "email": "test@email.com",
    "password": "9501611ee51cdc4dc3386c9788f72b273abb3325f951383a8e4b6c34a85712ac",
    "typ": 0,
    "avatar": null
}
*/
func (c *UserController) info(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	var res entity.User
	if err := repo.DB.First(&res, userId).Error; err != nil {
		ErrSys(ctx, err)
		return
	}

	ctx.JSON(200, res)
}

/**
@api {GET} /api/user/avatar 上传头像
@apiDescription 上传头像。
@apiName UserPostAvatar
@apiGroup User

@apiHeader {String} token 认证令牌
@apiHeader {String} Content-Type form/data

@apiSuccessExample {HTTP} 成功响应
HTTP/1.1 200 OK

*/
func (c *UserController) uploadAvatar(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	file, err := ctx.FormFile("img")
	if err != nil {
		ErrIllegal(ctx, "上传图片文件错误")
		return
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpeg" && ext != ".jpg" && ext != ".png" {
		ErrIllegal(ctx, "仅支持PNG或jpg格式")
		return
	}
	if file.Size >= 64*KB {
		ErrIllegal(ctx, "头像应小于64KB")
		return
	}
	ff, err := file.Open()
	if err != nil {
		ErrSys(ctx, err)
		return
	}
	defer ff.Close()
	avatarData, err := ioutil.ReadAll(ff)
	if err != nil {
		ErrSys(ctx, err)
		return
	}

	err = repo.DB.Model(&entity.User{}).Where("id = ?", userId).Update("avatar", avatarData).Error
	if err != nil {
		ErrSys(ctx, err)
		return
	}
}
