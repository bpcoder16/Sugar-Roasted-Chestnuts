package main

import (
	"Sugar-Roasted-Chestnuts/bootstrap"
	"Sugar-Roasted-Chestnuts/cmd/services"
	"context"
	"github.com/bpcoder16/Chestnut/core/cmd"
	"github.com/bpcoder16/Chestnut/modules/appconfig"
)

func main() {
	config := appconfig.MustLoadAppConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootstrap.MustInit(ctx, config)

	cmd.InitRootCmd(ctx)
	registerServices()
	cmd.Run()
}

func registerServices() {
	// 测试 Chestnut 基础库对应的 logit 模块
	cmd.RegisterService(&services.LogTest{})
	// 测试 Chestnut 基础库对应的默认 DefaultMySQL 模块
	cmd.RegisterService(&services.DefaultMySQL{})
	// 测试 Chestnut 基础库对应的多 MySQL 实例功能
	cmd.RegisterService(&services.MultipleMySQL{})
	// 测试 Chestnut 基础库对应的默认 DefaultRedis 模块
	cmd.RegisterService(&services.DefaultRedis{})
	// 测试 Chestnut 基础库对应的多 Redis 实例功能
	cmd.RegisterService(&services.MultipleRedis{})
	// 测试 Chestnut 基础库对应的默认 DefaultClickhouse 模块
	cmd.RegisterService(&services.DefaultClickhouse{})
	// 测试 Chestnut 基础库对应的多 Clickhouse 实例功能
	cmd.RegisterService(&services.MultipleClickhouse{})
	// 测试 Chestnut 基础库对应的默认 DefaultMongoDB 模块
	cmd.RegisterService(&services.DefaultMongoDB{})
	// 测试 Chestnut 基础库对应的多 MongoDB 实例功能
	cmd.RegisterService(&services.MultipleMongoDB{})
}
