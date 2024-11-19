package mysql

import (
	"github.com/bpcoder16/Chestnut/contrib/orm/mysql"
	"github.com/bpcoder16/Chestnut/core/log"
	"gorm.io/gorm"
)

var defaultMySQLTest01GormDBManager *mysql.GormDBManager
var defaultMySQLTest02GormDBManager *mysql.GormDBManager

func SetTest01Manager(configPath string, logger *log.Helper) {
	defaultMySQLTest01GormDBManager = mysql.NewGormDBManager(configPath, logger)
}

func Test01MasterDB() *gorm.DB {
	return defaultMySQLTest01GormDBManager.MasterDB()
}

func Test01SlaveDB() *gorm.DB {
	return defaultMySQLTest01GormDBManager.SlaveDB()
}

func SetTest02Manager(configPath string, logger *log.Helper) {
	defaultMySQLTest02GormDBManager = mysql.NewGormDBManager(configPath, logger)
}

func Test02MasterDB() *gorm.DB {
	return defaultMySQLTest02GormDBManager.MasterDB()
}

func Test02SlaveDB() *gorm.DB {
	return defaultMySQLTest02GormDBManager.SlaveDB()
}
