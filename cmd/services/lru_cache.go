package services

import (
	"context"
	"github.com/bpcoder16/Chestnut/cmd"
	"github.com/bpcoder16/Chestnut/logit"
	"github.com/bpcoder16/Chestnut/lru"
	"strconv"
	"time"
)

type LRUCache struct {
	cmd.Base
}

func (l *LRUCache) Description() string {
	return "测试 Chestnut 基础库对应的 LRUCache 模块"
}

func (l *LRUCache) Run(ctx context.Context, _ []string) {
	for i := 0; i < 10; i++ {
		lru.DefaultLRUCache().Add(strconv.Itoa(i), i)
	}
	logit.Context(ctx).DebugW("LRUCache.Len", lru.DefaultLRUCache().Len())
}

type LRUCacheExpire struct {
	cmd.Base
}

func (l *LRUCacheExpire) Description() string {
	return "测试 Chestnut 基础库对应的 LRUCacheExpire 模块"
}

func (l *LRUCacheExpire) Run(ctx context.Context, _ []string) {
	for i := 0; i < 10; i++ {
		lru.DefaultExpireLRUCache().Add(strconv.Itoa(i), i)
	}
	logit.Context(ctx).DebugW("LRUCache.Len", lru.DefaultExpireLRUCache().Len())
	time.Sleep(time.Minute)
}
