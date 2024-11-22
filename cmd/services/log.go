package services

import (
	"Sugar-Roasted-Chestnuts/services/logtest"
	"context"
	"github.com/bpcoder16/Chestnut/modules/cmd"
)

type LogTest struct {
	cmd.Base
}

func (l *LogTest) Description() string {
	return "测试 Chestnut 基础库对应的 logit 模块"
}

func (l *LogTest) Run(ctx context.Context, _ []string) {
	logtest.Test(ctx)
}
