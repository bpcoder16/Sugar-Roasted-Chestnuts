package redis

import (
	"github.com/bpcoder16/Chestnut/contrib/goredis"
	"github.com/bpcoder16/Chestnut/core/log"
	"github.com/redis/go-redis/v9"
)

var defaultRedisTest01Manager *goredis.Manager
var defaultRedisTest02Manager *goredis.Manager

func SetTest01Manager(configPath string, logger *log.Helper) {
	defaultRedisTest01Manager = goredis.NewManager(configPath, logger)
}

func Test01Client() *redis.Client {
	return defaultRedisTest01Manager.Client()
}

func SetTest02Manager(configPath string, logger *log.Helper) {
	defaultRedisTest02Manager = goredis.NewManager(configPath, logger)
}

func Test02Client() *redis.Client {
	return defaultRedisTest02Manager.Client()
}
