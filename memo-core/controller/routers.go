package controller

import (
	"github.com/gin-gonic/gin"
)

// routeMapping HTTP路由注册
func routeMapping(r gin.IRouter) {
	NewUserController(r)
}