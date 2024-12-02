package main

import (
	"Sugar-Roasted-Chestnuts/bootstrap"
	"Sugar-Roasted-Chestnuts/httpserver/route"
	"Sugar-Roasted-Chestnuts/pubsub"
	"context"
	"github.com/bpcoder16/Chestnut/appconfig"
	"github.com/bpcoder16/Chestnut/appconfig/env"
	"github.com/bpcoder16/Chestnut/contrib/httphandler/gin"
	"github.com/bpcoder16/Chestnut/core/cdefer"
	"github.com/bpcoder16/Chestnut/core/gtask"
	"github.com/bpcoder16/Chestnut/modules/httpserver"
	"path"
)

func main() {
	config := appconfig.MustLoadAppConfig("/conf/app.json")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootstrap.MustInit(ctx, config)
	defer cdefer.Defer()

	routeWebSocket := route.WebSocket()

	var g *gtask.Group
	g, ctx = gtask.WithContext(ctx)
	g.Go(func() error {
		httpserver.NewManager(
			path.Join(env.RootPath(), "conf/http.json"),
			gin.HTTPHandler(
				route.Api(),
				route.ApiV2(),
				routeWebSocket,
			),
		).Run()
		return nil
	})
	g.Go(func() error {
		return pubsub.RedisSubscribe(ctx, routeWebSocket.GetClientManager())
	})

	_ = g.Wait()
}
