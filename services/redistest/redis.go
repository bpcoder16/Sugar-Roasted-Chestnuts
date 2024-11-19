package redistest

import (
	SRedis "Sugar-Roasted-Chestnuts/redis"
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

func MultipleTest(ctx context.Context) {
	SRedis.Test01Client().Set(ctx, "demo01", "demo01", time.Second*5)

	result, err := SRedis.Test01Client().Get(ctx, "demo01").Result()
	logit.Context(ctx).InfoW("SRedis.Test01Client().result", result, "err", err)

	time.Sleep(time.Second * 5)

	result, err = SRedis.Test01Client().Get(ctx, "demo01").Result()
	logit.Context(ctx).InfoW("SRedis.Test01Client().result", result, "err", err)
	if errors.Is(err, goRedis.Nil) {
		logit.Context(ctx).InfoW("SRedis.Test01Client().result.Err", "goRedis.Nil")
	}

	SRedis.Test02Client().Set(ctx, "demo02", "demo02", time.Second*5)

	result, err = SRedis.Test02Client().Get(ctx, "demo02").Result()
	logit.Context(ctx).InfoW("SRedis.Test02Client().result", result, "err", err)

	time.Sleep(time.Second * 5)

	result, err = SRedis.Test02Client().Get(ctx, "demo02").Result()
	logit.Context(ctx).InfoW("SRedis.Test02Client().result", result, "err", err)
	if errors.Is(err, goRedis.Nil) {
		logit.Context(ctx).InfoW("SRedis.Test02Client().result.Err", "goRedis.Nil")
	}
}
