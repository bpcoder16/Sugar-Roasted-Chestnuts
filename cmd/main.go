package main

import (
	"Sugar-Roasted-Chestnuts/bootstrap"
	"Sugar-Roasted-Chestnuts/cmd/services"
	"context"
	"github.com/bpcoder16/Chestnut/appconfig"
	"github.com/bpcoder16/Chestnut/cmd"
	"github.com/bpcoder16/Chestnut/core/cdefer"
)

func main() {
	config := appconfig.MustLoadAppConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootstrap.MustInit(ctx, config)
	defer cdefer.Defer()

	cmd.InitRootCmd(ctx)
	services.RegisterServices()
	cmd.Run()
}
