package controller

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/tjfoc/gmsm/sm3"
	"memo-core/controller/midd"
	"memo-core/repo"
	"memo-core/repo/entity"
	"net/mail"
	"strconv"
)

func NewUserController(r gin.IRouter) *UserController {
	res := &UserController{}
	base := r.Group("user")
	// 创建用户
	base.POST("register", res.register)
	base.DELETE("", midd.AdminOnly, res.delete)
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
