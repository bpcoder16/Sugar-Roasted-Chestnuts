package mysql

import (
	"github.com/bpcoder16/Chestnut/contrib/orm/mysql"
	"github.com/bpcoder16/Chestnut/core/log"
	"gorm.io/gorm"
)

var defaultMySQLTest01GormDBManager *mysql.Manager
var defaultMySQLTest02GormDBManager *mysql.Manager

func SetTest01Manager(configPath string, logger *log.Helper) {
	defaultMySQLTest01GormDBManager = mysql.NewManager(configPath, logger)
}

func Test01MasterDB() *gorm.DB {
	return defaultMySQLTest01GormDBManager.MasterDB()
}

func Test01SlaveDB() *gorm.DB {
	return defaultMySQLTest01GormDBManager.SlaveDB()
}

func SetTest02Manager(configPath string, logger *log.Helper) {
	defaultMySQLTest02GormDBManager = mysql.NewManager(configPath, logger)
}

func Test02MasterDB() *gorm.DB {
	return defaultMySQLTest02GormDBManager.MasterDB()
}

func Test02SlaveDB() *gorm.DB {
	return defaultMySQLTest02GormDBManager.SlaveDB()
}
