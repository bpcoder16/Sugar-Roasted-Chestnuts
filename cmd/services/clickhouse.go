package services

import (
	"Sugar-Roasted-Chestnuts/services/clickhousetest"
	"context"
	"github.com/bpcoder16/Chestnut/modules/cmd"
)

type DefaultClickhouse struct {
	cmd.Base
}

func (d *DefaultClickhouse) Description() string {
	return "测试 Chestnut 基础库对应的默认 DefaultClickhouse 模块"
}

func (d *DefaultClickhouse) Run(ctx context.Context, _ []string) {
	clickhousetest.Test(ctx)
}

type MultipleClickhouse struct {
	cmd.Base
}

func (m *MultipleClickhouse) Description() string {
	return "测试 Chestnut 基础库对应的多 Clickhouse 实例功能"
}

func (m *MultipleClickhouse) Run(ctx context.Context, _ []string) {
	clickhousetest.MultipleTest(ctx)
}
