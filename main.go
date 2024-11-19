package main

import (
	"Sugar-Roasted-Chestnuts/bootstrap"
	"Sugar-Roasted-Chestnuts/services/redistest"
	"context"
	"github.com/bpcoder16/Chestnut/modules/appconfig"
)

func main() {
	config := appconfig.MustLoadAppConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//configtest.Test(ctx, config)

	bootstrap.MustInit(ctx, config)

	//logtest.Test(ctx)

	//mysqltest.Test(ctx)

	redistest.Test(ctx)
}
