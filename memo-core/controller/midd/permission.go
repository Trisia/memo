package midd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	TypUser  uint = 0
	TypAdmin uint = 1
	TypApp   uint = 2
	TypNone  uint = 0xFF
)

var (
	UserOnly  = BuildPermissionMid(TypUser)
	AdminOnly = BuildPermissionMid(TypAdmin)
	AppOnly   = BuildPermissionMid(TypApp)
	Entity    = BuildPermissionMid(TypUser, TypApp) // 实体包括 用户、应用系统
)

// BuildPermissionMid 构造权限认证中间件
func BuildPermissionMid(types ...uint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		typ := ctx.GetUint("typ")
		for _, u := range types {
			if typ == u {
				return
			}
		}
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
}
