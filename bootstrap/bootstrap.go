package bootstrap

import (
	"Sugar-Roasted-Chestnuts/mysql"
	"context"
	"github.com/bpcoder16/Chestnut/bootstrap"
	"github.com/bpcoder16/Chestnut/core/log"
	"github.com/bpcoder16/Chestnut/modules/appconfig"
	"github.com/bpcoder16/Chestnut/modules/appconfig/env"
	"github.com/bpcoder16/Chestnut/modules/zaplogger"
	"io"
)

func MustInit(ctx context.Context, config *appconfig.AppConfig) {
	bootstrap.MustInit(ctx, config, initMultipleMySQL)
}

func initMultipleMySQL(ctx context.Context, debugWriter, infoWriter, warnErrorFatalWriter io.Writer) {
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
