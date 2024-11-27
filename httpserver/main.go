package main

import (
	"Sugar-Roasted-Chestnuts/bootstrap"
	"Sugar-Roasted-Chestnuts/httpserver/route"
	"context"
	"github.com/bpcoder16/Chestnut/appconfig"
	"github.com/bpcoder16/Chestnut/appconfig/env"
	"github.com/bpcoder16/Chestnut/contrib/httphandler/gin"
	"github.com/bpcoder16/Chestnut/core/cdefer"
	"github.com/bpcoder16/Chestnut/modules/httpserver"
	"path"
)

func main() {
	config := appconfig.MustLoadAppConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootstrap.MustInit(ctx, config)
	defer cdefer.Defer()

	httpserver.NewManager(
		path.Join(env.RootPath(), "conf/http.json"),
		gin.HTTPHandler(
			route.Api(),
			route.ApiV2(),
			route.WebSocket(),
		),
	).Run()
}
