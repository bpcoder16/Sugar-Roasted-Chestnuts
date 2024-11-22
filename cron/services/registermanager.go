package services

import "github.com/bpcoder16/Chestnut/core/cron"

func RegisterServices() {
	//cron.RegisterCron("Demo1", &Demo{})
	//cron.RegisterCron("Demo2", &Demo{})
	cron.RegisterCron("Demo3", &Demo{})
}
