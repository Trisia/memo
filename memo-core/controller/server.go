package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"memo-core/controller/midd"
	"memo-core/setting"
	"net/http"
)

func NewServer() *http.Server {
	r := gin.Default()
	r.Use(
		Recovery(),
		midd.ValidateJWT,
	)
	r.MaxMultipartMemory = 64 << 20 // 64 MiB
	// 注册路路由
	routeMapping(r)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Config.Port),
		Handler: r,
	}
}
