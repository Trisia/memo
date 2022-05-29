package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"memo-core/repo"
	"memo-core/setting"
)

func main() {
	err := setting.Load()
	if err != nil {
		zap.L().Fatal("配置初始化异常", zap.Error(err))
	}
	err = repo.Init(setting.Config.Database)
	if err != nil {
		zap.L().Fatal("数据持久化服务启动异常", zap.Error(err))
	}

	server := gin.Default()
	zap.L().Info("启动memo-core服务", zap.Int("port", setting.Config.Port))
	err = server.Run(fmt.Sprintf(":%d", setting.Config.Port))
	if err != nil {
		zap.L().Fatal("服务启动异常", zap.Error(err))
	}
}
