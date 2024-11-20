package main

import (
	"Sugar-Roasted-Chestnuts/bootstrap"
	"Sugar-Roasted-Chestnuts/services/aliyun"
	"context"
	"github.com/bpcoder16/Chestnut/modules/appconfig"
)

func main() {
	config := appconfig.MustLoadAppConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootstrap.MustInit(ctx, config)

	//clickhousetest.Test(ctx)
	//clickhousetest.MultipleTest(ctx)

	//mongodbtest.Test(ctx)
	//mongodbtest.MultipleTest(ctx)

	aliyun.ImageTransferTest(ctx)
}
