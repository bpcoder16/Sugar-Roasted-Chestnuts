package redistest

import (
	"context"
	"errors"
	"github.com/bpcoder16/Chestnut/logit"
	"github.com/bpcoder16/Chestnut/redis"
	goRedis "github.com/redis/go-redis/v9"
	"time"
)

func Test(ctx context.Context) {
	redis.DefaultClient().Set(ctx, "demo01", "demo01", time.Second*5)

	result, err := redis.DefaultClient().Get(ctx, "demo01").Result()
	logit.Context(ctx).InfoW("result", result, "err", err)

	time.Sleep(time.Second * 5)

	result, err = redis.DefaultClient().Get(ctx, "demo01").Result()
	logit.Context(ctx).InfoW("result", result, "err", err)
	if errors.Is(err, goRedis.Nil) {
		logit.Context(ctx).InfoW("result.Err", "goRedis.Nil")
	}
}
