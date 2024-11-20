package services

import (
	"Sugar-Roasted-Chestnuts/services/mysqltest"
	"context"
	"github.com/bpcoder16/Chestnut/core/cmd"
)

type DefaultMySQL struct {
	cmd.Base
}

func (d *DefaultMySQL) Description() string {
	return "测试 Chestnut 基础库对应的默认 DefaultMySQL 模块"
}

func (d *DefaultMySQL) Run(ctx context.Context, _ []string) {
	mysqltest.Test(ctx)
}

type MultipleMySQL struct {
	cmd.Base
}

func (m *MultipleMySQL) Description() string {
	return "测试 Chestnut 基础库对应的多 MySQL 实例功能"
}

func (m *MultipleMySQL) Run(ctx context.Context, _ []string) {
	mysqltest.MultipleTest(ctx)
}
