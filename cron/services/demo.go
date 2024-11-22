package services

import (
	"github.com/bpcoder16/Chestnut/logit"
	"github.com/bpcoder16/Chestnut/modules/cron"
)

type Demo struct {
	cron.Base
}

func (d *Demo) Process() {
	logit.Context(d.Ctx).DebugW("info", "Demo Running.")
}
