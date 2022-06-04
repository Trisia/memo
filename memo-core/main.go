package main

import (
	"go.uber.org/zap"
	"memo-core/controller"
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

	zap.L().Info("启动memo-core服务", zap.Int("port", setting.Config.Port))
	server := controller.NewServer()
	err = server.ListenAndServe()
	if err != nil {
		zap.L().Fatal("HTTP服务异常终止", zap.Error(err))
	}
}
