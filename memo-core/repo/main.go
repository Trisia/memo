package repo

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"memo-core/setting"
	"strings"
	"time"
)

var DB *gorm.DB

var (
	UserSvc      *UserRepo
	DocTagAggSvc *DocTagAggRepo
)

func register() {
	UserSvc = NewUserRepo()
	DocTagAggSvc = NewDocTagAggRepo()
}

// Init 初始化数据库连接服务
func Init(config setting.Database) error {
	if len(config.Dsn) == 0 {
		return fmt.Errorf("数据库系统 DSN为空")
	}
	config.Type = strings.ToUpper(config.Type)
	var dialector gorm.Dialector
	switch config.Type {
	case "MYSQL":
		dialector = mysql.Open(config.Dsn)
	case "SQLITE", "":
		dialector = sqlite.Open(config.Dsn)
	default:
		return fmt.Errorf("未知的数据库系统类型 [%s]", config.Type)
	}
	var err error
	opts := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	if setting.Config != nil && setting.Config.Debug {
		// 打印DEBUG日志
		opts.Logger = logger.Default.LogMode(logger.Info)
	}

	DB, err = gorm.Open(dialector, opts)
	if err != nil {
		return fmt.Errorf("无法连接数据库, %v", err)
	}
	sqlDB, err := DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(2)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(128)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	zap.L().Info(fmt.Sprintf("数据库 [%s] 连接成功", config.Type))
	// 创建数据库表
	err = Migrate(false)
	if err != nil {
		return err
	}

	// 注册仓库
	register()
	return nil
}
