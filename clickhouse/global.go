package clickhouse

import (
	"github.com/bpcoder16/Chestnut/contrib/orm/clickhouse"
	"github.com/bpcoder16/Chestnut/core/log"
	"gorm.io/gorm"
)

var defaultClickhouseTest01GormDBManager *clickhouse.Manager
var defaultClickhouseTest02GormDBManager *clickhouse.Manager

func SetTest01Manager(configPath string, logger *log.Helper) {
	defaultClickhouseTest01GormDBManager = clickhouse.NewManager(configPath, logger)
}

func Test01MasterDB() *gorm.DB {
	return defaultClickhouseTest01GormDBManager.MasterDB()
}

func Test01SlaveDB() *gorm.DB {
	return defaultClickhouseTest01GormDBManager.SlaveDB()
}

func SetTest02Manager(configPath string, logger *log.Helper) {
	defaultClickhouseTest02GormDBManager = clickhouse.NewManager(configPath, logger)
}

func Test02MasterDB() *gorm.DB {
	return defaultClickhouseTest02GormDBManager.MasterDB()
}

func Test02SlaveDB() *gorm.DB {
	return defaultClickhouseTest02GormDBManager.SlaveDB()
}
