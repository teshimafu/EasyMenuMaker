package migrations

import (
	"github.com/teshimafu/lazyPM/server/src/infrastructure/model"
	"gorm.io/gorm"
)

// マイグレーションの実行
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	return nil
}
