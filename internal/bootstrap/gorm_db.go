package bootstrap

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"memo/internal/configs"
)

func NewGORMDb(conf configs.DBConfig) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(conf.Database), &gorm.Config{})
}
