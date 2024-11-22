package services

import (
	"Sugar-Roasted-Chestnuts/services/clickhousetest"
	"Sugar-Roasted-Chestnuts/services/logtest"
	"Sugar-Roasted-Chestnuts/services/mongodbtest"
	"Sugar-Roasted-Chestnuts/services/mysqltest"
	"Sugar-Roasted-Chestnuts/services/redistest"
	"github.com/bpcoder16/Chestnut/cron"
)

type Demo struct {
	cron.Base
}

func (d *Demo) Process() {
	logtest.Test(d.Ctx)
	mysqltest.Test(d.Ctx)
	redistest.Test(d.Ctx)
	mongodbtest.Test(d.Ctx)
	clickhousetest.Test(d.Ctx)
}
