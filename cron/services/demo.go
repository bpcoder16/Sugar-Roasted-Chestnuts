package services

import (
	"context"
	"github.com/bpcoder16/Chestnut/cron"
	"github.com/bpcoder16/Chestnut/logit"
)

type Demo struct {
	Base
	Desc string
}

func (d *Demo) Init(s cron.Interface) {
	d.Base.Init(s)
	if b, ok := s.(*Demo); ok && b != nil {
		d.Desc = b.Desc
	}
}

func (d *Demo) Process() {
	for i := 0; i < 2; i++ {
		d.AddProcessAddTaskList(func(ctx context.Context) {
			logit.Context(ctx).InfoW("xxx", "xxx", "desc", d.Desc)
		})
	}
	//logtest.Test(d.Ctx)
	//mysqltest.Test(d.Ctx)
	//redistest.Test(d.Ctx)
	//mongodbtest.Test(d.Ctx)
	//clickhousetest.Test(d.Ctx)
}
