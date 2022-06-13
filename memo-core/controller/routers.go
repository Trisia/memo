package controller

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"memo-core/setting"
	"mime"
)

// routeMapping HTTP路由注册
func routeMapping(r gin.IRouter) {
	_ = mime.AddExtensionType(".js", "application/javascript")
	r.Use(static.Serve("/", static.LocalFile(setting.LocWeb, true)))
	r = r.Group("/api")
	NewAuthController(r)
	NewUserController(r)
	NewDocController(r)
	NewTagController(r)
}
