package setting

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

// logInit 初始化日志
func logInit() {
	if Logger != nil {
		return
	}

	// 创建文件目录
	fileLogger := &lumberjack.Logger{
		Filename:  LocLog,
		MaxSize:   3,  // MB
		MaxAge:    28, // Day
		LocalTime: true,
		Compress:  true,
	}
	productWriter := zapcore.AddSync(fileLogger)

	pe := zap.NewProductionEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(pe)
	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	// 同时向控制台和文件写入日志
	zapLevel := zapcore.InfoLevel
	ginLevel := gin.ReleaseMode

	if Config.Debug {
		zapLevel = zapcore.DebugLevel
		ginLevel = gin.DebugMode
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, productWriter, zapLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapLevel),
	)
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	gin.SetMode(ginLevel)
	Logger = logger
}
