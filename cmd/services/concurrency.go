package services

import (
	"context"
	"github.com/bpcoder16/Chestnut/cmd"
	"github.com/bpcoder16/Chestnut/logit"
	"github.com/bpcoder16/Chestnut/modules/concurrency"
	"strconv"
)

type ConcurrencyTest struct {
	cmd.Base
}

func (c *ConcurrencyTest) Description() string {
	return "测试 Chestnut 基础库对应的 Concurrency 模块"
}

func (c *ConcurrencyTest) Run(ctx context.Context, _ []string) {
	taskMap := make(map[string]func(ctx context.Context) concurrency.ChanResult)
	for i := 0; i < 2; i++ {
		taskMap[strconv.Itoa(i)] = func(ctxTask context.Context) concurrency.ChanResult {
			logit.Context(ctxTask).InfoW("xxx", "xxx")
			return concurrency.ChanResult{}
		}
	}
	resultMap, err := concurrency.Manager(ctx, taskMap, "ConcurrencyTest")
	logit.Context(ctx).InfoW("resultMap", resultMap, "err", err)
}
