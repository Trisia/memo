package controller

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/tjfoc/gmsm/sm3"
	"memo-core/repo"
	"memo-core/repo/entity"
	"net/mail"
)

func NewUserController(r gin.IRouter) *UserController {
	res := &UserController{}
	base := r.Group("user")
	base.POST("create", res.create)
	return res
}

// UserController 用户控制器
type UserController struct {
}

/**
@api {GET} /user/create 创建用户
@apiDescription 创建用户。
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

@apiSuccessExample {json} 成功响应

*/
func (c *UserController) create(ctx *gin.Context) {
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
