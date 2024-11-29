package services

import (
	"github.com/bpcoder16/Chestnut/cron"
)

func RegisterServices() {
	cron.RegisterCron("Demo1", &Demo{Desc: "Desc"})
	//cron.RegisterCron("Demo2", &Demo{})
	//cron.RegisterCron("Demo3", &Demo{})
}
