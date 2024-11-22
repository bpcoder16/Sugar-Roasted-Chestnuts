package services

import (
	"context"
	"github.com/bpcoder16/Chestnut/cron"
	"github.com/bpcoder16/Chestnut/logit"
)

type Demo struct {
	cron.Base
}

func (d *Demo) Process() {
	for i := 0; i < 2; i++ {
		d.AddProcessAddTaskList(func(ctx context.Context) {
			logit.Context(ctx).InfoW("xxx", "xxx")
		})
	}
	//logtest.Test(d.Ctx)
	//mysqltest.Test(d.Ctx)
	//redistest.Test(d.Ctx)
	//mongodbtest.Test(d.Ctx)
	//clickhousetest.Test(d.Ctx)
}
