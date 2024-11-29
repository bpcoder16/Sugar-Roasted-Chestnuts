package services

import (
	"context"
	"github.com/bpcoder16/Chestnut/cmd"
)

type Demo struct {
	cmd.Base
}

func (demo *Demo) Description() string {
	return "测试各种临时功能代码"
}

func (demo *Demo) Run(_ context.Context, _ []string) {}
