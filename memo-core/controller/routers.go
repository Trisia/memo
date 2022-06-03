package controller

import (
	"github.com/gin-gonic/gin"
)

// RouteMapping HTTP路由注册
func RouteMapping(r gin.IRouter) {
	NewUserController(r)
}
