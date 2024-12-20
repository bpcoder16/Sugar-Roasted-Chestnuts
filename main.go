package main

import (
	"Sugar-Roasted-Chestnuts/bootstrap"
	"context"
	"github.com/bpcoder16/Chestnut/appconfig"
)

func main() {
	config := appconfig.MustLoadAppConfig("/conf/app.json")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootstrap.MustInit(ctx, config)
}
