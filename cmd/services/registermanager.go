package services

import "github.com/bpcoder16/Chestnut/cmd"

func RegisterServices() {
	// 测试 Chestnut 基础库对应的 logit 模块
	cmd.RegisterService(&LogTest{})
	// 测试 Chestnut 基础库对应的默认 DefaultMySQL 模块
	cmd.RegisterService(&DefaultMySQL{})
	// 测试 Chestnut 基础库对应的多 MySQL 实例功能
	cmd.RegisterService(&MultipleMySQL{})
	// 测试 Chestnut 基础库对应的默认 DefaultRedis 模块
	cmd.RegisterService(&DefaultRedis{})
	// 测试 Chestnut 基础库对应的多 Redis 实例功能
	cmd.RegisterService(&MultipleRedis{})
	// 测试 Chestnut 基础库对应的默认 DefaultClickhouse 模块
	cmd.RegisterService(&DefaultClickhouse{})
	// 测试 Chestnut 基础库对应的多 Clickhouse 实例功能
	cmd.RegisterService(&MultipleClickhouse{})
	// 测试 Chestnut 基础库对应的默认 DefaultMongoDB 模块
	cmd.RegisterService(&DefaultMongoDB{})
	// 测试 Chestnut 基础库对应的多 MongoDB 实例功能
	cmd.RegisterService(&MultipleMongoDB{})
	// 测试 Chestnut 基础库对应的 ImageTransfer 模块
	cmd.RegisterService(&AliyunOSSImageTransfer{})
	// 测试 Chestnut 基础库对应的 AliyunOSSSignURL 模块
	cmd.RegisterService(&AliyunOSSSignURL{})
	// 测试 Chestnut 基础库对应的 LRUCache 模块
	cmd.RegisterService(&LRUCache{})
	// 测试 Chestnut 基础库对应的 LRUCacheExpire 模块
	cmd.RegisterService(&LRUCacheExpire{})
	// 测试 Chestnut 基础库对应的 Concurrency 模块
	cmd.RegisterService(&ConcurrencyTest{})
	// 测试各种临时功能代码
	cmd.RegisterService(&Demo{})
}
