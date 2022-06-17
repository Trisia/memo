package midd

import (
	"github.com/gin-gonic/gin"
	"memo-core/controller/jwt"
	"net/http"
	"strconv"
)

// ValidateJWT 验证JWT有效性
// 校验Token 提取信息到上下文中
func ValidateJWT(c *gin.Context) {
	c.Set("typ", uint(0xFF))
	// 从请求头中获取认证参数
	token := c.GetHeader("token")
	if token == "" {
		return
	}
	claims := jwt.Verify(token)
	if claims == nil {
		// Token 无效，禁止访问
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId, _ := strconv.Atoi(claims.Sub)
	// 设置上下文
	c.Set("typ", claims.Typ)
	c.Set("userId", uint(userId))
}
