package services

import (
	"github.com/bpcoder16/Chestnut/cron"
	"github.com/bpcoder16/Chestnut/redis"
)

type Base struct {
	cron.Base
}

func (b *Base) Init(s cron.Interface) {
	b.RedisClient = redis.DefaultClient()
	b.Base.Init(s)
}
