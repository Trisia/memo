package setting

import (
	"os"
	"path/filepath"
)

var (
	LocExe    string // 可执行文件路径
	LocBase   string // 基本文件路径
	LocConfig string // 配置文件路径
	LocLog    string // 日志文件路径
	LocWeb    string // web文件路径
)

// 初始化文件路径
func filepathInit() {
	LocExe, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	// 日志文件目录
	_ = os.MkdirAll(filepath.Join(LocBase, "logs/"), os.ModePerm)
	LocLog = filepath.Join(LocBase, "logs", "memo.log")
	LocConfig = filepath.Join(LocBase, "application.yaml")
	LocWeb = filepath.Join(LocExe, "web/")
	_ = os.MkdirAll(LocWeb, os.ModePerm)
}
