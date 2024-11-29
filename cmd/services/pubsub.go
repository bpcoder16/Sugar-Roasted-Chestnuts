package services

import (
	"Sugar-Roasted-Chestnuts/pubsub"
	"context"
	"github.com/bpcoder16/Chestnut/cmd"
	"github.com/bpcoder16/Chestnut/contrib/websocket"
)

type PubSub struct {
	cmd.Base
}

func (p *PubSub) Description() string {
	return "测试 Chestnut 基础库对应的 PubSub 模块"
}

func (p *PubSub) Run(ctx context.Context, _ []string) {
	pubsub.RedisPublish(ctx, pubsub.Message{
		Code: 0,
		State: websocket.State{
			Scene: "test",
		},
		ID:   "11111111111111",
		Data: "bbbbbbbb",
	})
}
