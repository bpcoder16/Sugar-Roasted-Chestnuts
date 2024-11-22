package services

import (
	"Sugar-Roasted-Chestnuts/services/mongodbtest"
	"context"
	"github.com/bpcoder16/Chestnut/cmd"
)

type DefaultMongoDB struct {
	cmd.Base
}

func (d *DefaultMongoDB) Description() string {
	return "测试 Chestnut 基础库对应的默认 DefaultMongoDB 模块"
}

func (d *DefaultMongoDB) Run(ctx context.Context, _ []string) {
	mongodbtest.Test(ctx)
}

type MultipleMongoDB struct {
	cmd.Base
}

func (m *MultipleMongoDB) Description() string {
	return "测试 Chestnut 基础库对应的多 MongoDB 实例功能"
}

func (m *MultipleMongoDB) Run(ctx context.Context, _ []string) {
	mongodbtest.MultipleTest(ctx)
}
