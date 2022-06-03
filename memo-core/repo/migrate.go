package repo

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm3"
	"go.uber.org/zap"
	"memo-core/repo/entity"
)

// Migrate 迁移并创建数据库表，仅在user表不存在是运行迁移
// force: true - 强制迁移
func Migrate(force bool) error {
	if DB.Migrator().HasTable(&entity.User{}) && !force {
		return nil
	}
	err := DB.Migrator().AutoMigrate(
		&entity.User{},
		&entity.Document{},
		&entity.Tag{},
		&entity.DocTag{})
	if err != nil {
		return err
	}
	tables, _ := DB.Migrator().GetTables()
	zap.L().Info("创建表",
		zap.String("database", DB.Migrator().CurrentDatabase()),
		zap.Strings("table", tables))
	// 初始化管理员用户
	salt := make([]byte, 16)
	_, _ = rand.Reader.Read(salt)

	hash := sm3.New()
	hash.Write([]byte("123qwe!!"))
	hash.Write(salt)
	password := hash.Sum(nil)
	_ = DB.Create(&entity.User{
		Username: "admin",
		Password: hex.EncodeToString(password),
		Salt:     hex.EncodeToString(salt),
		Typ:      1,
	})

	return nil
}
