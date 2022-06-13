package controller

import (
	"github.com/gin-gonic/gin"
	"memo-core/controller/jwt"
	"net/http"
	"strconv"
)

var allowPath = map[string]uint8{
	"/api/user/register": 1,
	"/api/auth":          1,
}

// 验证JWT有效性
func validateJWT(c *gin.Context) {
	// 匿名接口
	if anonymous(c) {
		return
	}

	// 从请求头中获取认证参数
	token := c.GetHeader("token")

	if token == "" {
		// Token 无效，禁止访问
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims := jwt.Verify(token, jwtKey)
	if claims == nil {
		// Token 无效，禁止访问
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// 设置上下文
	userId, _ := strconv.Atoi(claims.Sub)
	c.Set("userId", uint(userId))
	c.Set("typ", claims.Typ)
}

// 检查路由是否是匿名
func anonymous(c *gin.Context) bool {
	_, ok := allowPath[c.Request.URL.Path]
	return ok
}
