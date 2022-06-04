package controller

import (
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"memo-core/controller/jwt"
	"memo-core/setting"
	"net/http"
)

// JWT的HMAC密钥
var jwtKey = make([]byte, 16)

func init() {
	_, _ = rand.Reader.Read(jwtKey)
}

func NewServer() *http.Server {
	r := gin.Default()
	r.Use(
		Recovery(),
		validateJWT,
	)
	// 注册路路由
	routeMapping(r)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Config.Port),
		Handler: r,
	}
}

// 验证JWT有效性
func validateJWT(c *gin.Context) {
	// 从请求头中获取认证参数
	token := c.GetHeader("token")

	if token == "" {
		return
	}

	claims := jwt.Verify(token, jwtKey)
	if claims == nil {
		// Token 无效，禁止访问
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// 设置上下文
	c.Set("username", claims.Sub)
	c.Set("typ", claims.Typ)
}
