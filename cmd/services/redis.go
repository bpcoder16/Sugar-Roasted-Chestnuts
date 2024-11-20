package services

import (
	"Sugar-Roasted-Chestnuts/services/redistest"
	"context"
	"github.com/bpcoder16/Chestnut/core/cmd"
)

type DefaultRedis struct {
	cmd.Base
}

func (d *DefaultRedis) Description() string {
	return "测试 Chestnut 基础库对应的默认 DefaultRedis 模块"
}

func (d *DefaultRedis) Run(ctx context.Context, _ []string) {
	redistest.Test(ctx)
}

type MultipleRedis struct {
	cmd.Base
}

func (m *MultipleRedis) Description() string {
	return "测试 Chestnut 基础库对应的多 Redis 实例功能"
}

func (m *MultipleRedis) Run(ctx context.Context, _ []string) {
	redistest.MultipleTest(ctx)
}
