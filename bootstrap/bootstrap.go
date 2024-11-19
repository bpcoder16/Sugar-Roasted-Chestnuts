package bootstrap

import (
	"Sugar-Roasted-Chestnuts/clickhouse"
	"Sugar-Roasted-Chestnuts/mysql"
	"Sugar-Roasted-Chestnuts/redis"
	"context"
	"github.com/bpcoder16/Chestnut/bootstrap"
	"github.com/bpcoder16/Chestnut/core/log"
	"github.com/bpcoder16/Chestnut/modules/appconfig"
	"github.com/bpcoder16/Chestnut/modules/appconfig/env"
	"github.com/bpcoder16/Chestnut/modules/zaplogger"
	"io"
)

func MustInit(ctx context.Context, config *appconfig.AppConfig) {
	bootstrap.MustInit(ctx, config)
	//bootstrap.MustInit(ctx, config, initMultipleMySQL, initMultipleRedis, initMultipleClickhouse)
}

func initMultipleMySQL(_ context.Context, debugWriter, infoWriter, warnErrorFatalWriter io.Writer) {
	mysql.SetTest01Manager(env.RootPath()+"/conf/test01_mysql.json", log.NewHelper(
		zaplogger.GetZapLogger(
			debugWriter, infoWriter, warnErrorFatalWriter,
			nil,
			log.FilterLevel(func() log.Level {
				if env.RunMode() == env.RunModeRelease {
					return log.LevelInfo
				}
				return log.LevelDebug
			}()),
			//log.FilterFunc(func(level log.Level, keyValues ...interface{}) bool {
			//	return false
			//}),
		),
	))
	mysql.SetTest02Manager(env.RootPath()+"/conf/test02_mysql.json", log.NewHelper(
		zaplogger.GetZapLogger(
			debugWriter, infoWriter, warnErrorFatalWriter,
			nil,
			log.FilterLevel(func() log.Level {
				if env.RunMode() == env.RunModeRelease {
					return log.LevelInfo
				}
				return log.LevelDebug
			}()),
			//log.FilterFunc(func(level log.Level, keyValues ...interface{}) bool {
			//	return false
			//}),
		),
	))
}

func initMultipleRedis(_ context.Context, debugWriter, infoWriter, warnErrorFatalWriter io.Writer) {
	redis.SetTest01Manager(env.RootPath()+"/conf/test01_redis.json", log.NewHelper(
		zaplogger.GetZapLogger(
			debugWriter, infoWriter, warnErrorFatalWriter,
			log.FileWithLineNumCallerRedis(),
			log.FilterLevel(func() log.Level {
				if env.RunMode() == env.RunModeRelease {
					return log.LevelInfo
				}
				return log.LevelDebug
			}()),
			//log.FilterFunc(func(level log.Level, keyValues ...interface{}) bool {
			//	return false
			//}),
		),
	))
	redis.SetTest02Manager(env.RootPath()+"/conf/test02_redis.json", log.NewHelper(
		zaplogger.GetZapLogger(
			debugWriter, infoWriter, warnErrorFatalWriter,
			log.FileWithLineNumCallerRedis(),
			log.FilterLevel(func() log.Level {
				if env.RunMode() == env.RunModeRelease {
					return log.LevelInfo
				}
				return log.LevelDebug
			}()),
			//log.FilterFunc(func(level log.Level, keyValues ...interface{}) bool {
			//	return false
			//}),
		),
	))
}

func initMultipleClickhouse(_ context.Context, debugWriter, infoWriter, warnErrorFatalWriter io.Writer) {
	clickhouse.SetTest01Manager(env.RootPath()+"/conf/test01_clickhouse.json", log.NewHelper(
		zaplogger.GetZapLogger(
			debugWriter, infoWriter, warnErrorFatalWriter,
			nil,
			log.FilterLevel(func() log.Level {
				if env.RunMode() == env.RunModeRelease {
					return log.LevelInfo
				}
				return log.LevelDebug
			}()),
			//log.FilterFunc(func(level log.Level, keyValues ...interface{}) bool {
			//	return false
			//}),
		),
	))
	clickhouse.SetTest02Manager(env.RootPath()+"/conf/test02_clickhouse.json", log.NewHelper(
		zaplogger.GetZapLogger(
			debugWriter, infoWriter, warnErrorFatalWriter,
			nil,
			log.FilterLevel(func() log.Level {
				if env.RunMode() == env.RunModeRelease {
					return log.LevelInfo
				}
				return log.LevelDebug
			}()),
			//log.FilterFunc(func(level log.Level, keyValues ...interface{}) bool {
			//	return false
			//}),
		),
	))
}
