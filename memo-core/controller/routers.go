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
	//_ = mime.AddExtensionType(".svg", "image/svg+xml")
	//_ = mime.AddExtensionType(".png", "image/png")
	//_ = mime.AddExtensionType(".jpeg", "image/jpeg")
	//_ = mime.AddExtensionType(".jpg", "image/jpeg")
	//_ = mime.AddExtensionType(".jpg", "image/jpeg")
	//_ = mime.AddExtensionType(".pdf", "application/pdf")

	r.Use(static.Serve("/", static.LocalFile(setting.LocWeb, true)))
	r = r.Group("/api")
	NewAuthController(r)
	NewUserController(r)
	NewDocController(r)
	NewDocAssetController(r)
	NewTagController(r)
}
