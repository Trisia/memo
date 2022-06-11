package controller

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tjfoc/gmsm/sm3"
	"gorm.io/gorm"
	"memo-core/controller/dto"
	"memo-core/controller/jwt"
	"memo-core/repo"
	"memo-core/repo/entity"
)

func NewAuthController(r gin.IRouter) *AuthController {
	res := &AuthController{}
	// 认证获取Token
	r.POST("auth", res.auth)
	return res
}

// AuthController 认证接口
type AuthController struct {
}

/**
@api {POST} /auth 认证
@apiDescription 用户认证获取Token。
@apiName Auth
@apiGroup Auth

@apiParam {String} username 用户名
@apiParam {String} password

@apiParamExample {json} 请求示例
{
	"username": "user1",
	"password": "123qwe"
}

@apiSuccess {String} Body token

@apiSuccessExample {String} 成功响应
eyJhbGciOiJITUFDLVNNMyIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0IiwidHlwIjowLCJpYXQiOjE2NTQ2OTk3MjExMTAsImV4cCI6MTY1NDcyODUyMTExMH0.OIva_EVx6Vy0PtJxmXxspM85M5QboDrU8T_BfC4meHY

@apiErrorExample {http} 失败
HTTP/1.1 400

用户名或口令错误
*/
func (c *AuthController) auth(ctx *gin.Context) {
	var param dto.UserAuth
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ErrIllegal(ctx, "无法解析参数")
		return
	}

	if param.Username == "" {
		ErrIllegal(ctx, "用户名为空")
		return
	}
	var user entity.User
	err := repo.DB.First(&user, "username = ?", param.Username).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ErrIllegal(ctx, "用户名或口令错误")
			return
		} else {
			ErrIllegal(ctx, "用户名为空")
			return
		}
	}
	salt, _ := hex.DecodeString(user.Salt)
	password, _ := hex.DecodeString(user.Password)
	hash := sm3.New()
	hash.Write([]byte(param.Password))
	hash.Write(salt)
	if bytes.Equal(hash.Sum(nil), password) {
		// 创建token
		token := jwt.Create(&jwt.Claims{Typ: user.Typ, Sub: fmt.Sprintf("%d", user.ID)}, jwtKey)
		ctx.String(200, token)
	} else {
		ErrIllegal(ctx, "用户名或口令错误")
		return
	}
}
