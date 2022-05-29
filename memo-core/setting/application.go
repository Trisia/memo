package setting

import "path/filepath"

// Application 配置对象
type Application struct {
	Database Database `yaml:"database"` // 数据库连接配置
	Port     int      `yaml:"port"`     // 服务端端口
	Debug    bool     `yaml:"debug"`    // 开发调试模式
}

// Database 数据库配置
type Database struct {
	Type string // 数据库类型：mysql、sqlite
	Dsn  string // 连接地址
}

// 无法找到配置文件时候的缺省配置
func newDefaultConfig() *Application {
	return &Application{
		//Database: Database{
		//	Dsn:   "user:pass@tcp(127.0.0.1:3306)/memo?charset=utf8mb4&parseTime=True&loc=Local",
		//	Type: "mysql",
		//},
		Database: Database{
			Dsn:  filepath.Join(LocExe, "memo.db"),
			Type: "sqlite",
		},
		Port:  7532,
		Debug: false,
	}
}
