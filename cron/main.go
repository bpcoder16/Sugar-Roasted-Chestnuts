package main

import (
	"Sugar-Roasted-Chestnuts/bootstrap"
	"Sugar-Roasted-Chestnuts/cron/services"
	"context"
	"github.com/bpcoder16/Chestnut/appconfig"
	"github.com/bpcoder16/Chestnut/appconfig/env"
	"github.com/bpcoder16/Chestnut/core/cdefer"
	"github.com/bpcoder16/Chestnut/cron"
	"path"
)

func main() {
	config := appconfig.MustLoadAppConfig("/conf/app.json")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootstrap.MustInit(ctx, config)
	defer cdefer.Defer()

	// 强依赖 Redis 支持
	services.RegisterServices()
	cron.Run(ctx, path.Join(env.ConfigDirPath(), "cron.json"))
}
