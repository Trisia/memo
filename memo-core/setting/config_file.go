package setting

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

// Config 配置参数 [只读]
var Config *Application

func init() {

	// 初始化文件路径
	filepathInit()
}

// Load 从配置文件中加在配置信息
func Load() error {
	bin, _ := ioutil.ReadFile(LocConfig)
	if len(bin) == 0 {
		// 没有加载到配置文件的情况使用默认配置
		Config = newDefaultConfig()
		// 生成默认配置文件
		b, _ := yaml.Marshal(Config)
		_ = ioutil.WriteFile(LocConfig, b, os.FileMode(0666))
		zap.L().Info("创建配置文件", zap.String("path", LocConfig))
	} else {
		zap.L().Info("读取配置文件", zap.String("path", LocConfig))
		Config = &Application{}
		err := yaml.Unmarshal(bin, Config)
		if err != nil {
			return err
		}
		if Config.Port <= 0 {
			Config.Port = 7743
		}
	}
	// 初始化日志配置
	logInit()

	return nil
}
