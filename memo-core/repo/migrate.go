package repo

import "memo-core/repo/entity"

// Migrate 迁移并创建数据库表，仅在user表不存在是运行迁移
// force: true - 强制迁移
func Migrate(force bool) error {
	if db.Migrator().HasTable(&entity.User{}) && !force {
		return nil
	}
	return db.Migrator().AutoMigrate(&entity.User{})
}
