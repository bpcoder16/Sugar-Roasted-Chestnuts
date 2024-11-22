package services

import (
	"github.com/bpcoder16/Chestnut/core/cron"
	"github.com/bpcoder16/Chestnut/logit"
)

type Demo struct {
	cron.Base
}

func (d *Demo) Process() {
	logit.Context(d.Ctx).DebugW("info", "Demo Running.")
}
