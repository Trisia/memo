package repo

import (
	"go.uber.org/zap"
	"memo-core/repo/entity"
)

// Migrate 迁移并创建数据库表，仅在user表不存在是运行迁移
// force: true - 强制迁移
func Migrate(force bool) error {
	if db.Migrator().HasTable(&entity.User{}) && !force {
		return nil
	}
	err := db.Migrator().AutoMigrate(
		&entity.User{},
		&entity.Document{},
		&entity.Tag{},
		&entity.DocTag{})
	tables, _ := db.Migrator().GetTables()
	zap.L().Info("schema创建完成",
		zap.String("database", db.Migrator().CurrentDatabase()),
		zap.Strings("table", tables))
	if err != nil {
		return err
	}
	return nil
}
