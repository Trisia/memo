package controller

import (
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
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
	r.MaxMultipartMemory = 64 << 20 // 64 MiB
	// 注册路路由
	routeMapping(r)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Config.Port),
		Handler: r,
	}
}
