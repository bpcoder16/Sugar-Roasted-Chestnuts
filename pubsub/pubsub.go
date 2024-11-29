package pubsub

import (
	"context"
	"encoding/json"
	"github.com/bpcoder16/Chestnut/appconfig/env"
	"github.com/bpcoder16/Chestnut/contrib/websocket"
	cRedis "github.com/bpcoder16/Chestnut/redis"
	"github.com/redis/go-redis/v9"
	"strconv"
	"sync"
)

const (
	pubSubChannelCnt = 2
	SceneTest        = "test"
)

type Message struct {
	Code            int `json:"code"`
	websocket.State `json:",inline"`
	ID              string `json:"id,omitempty"`   // 场景唯一 ID，如果场景没有唯一 ID，则空
	Data            any    `json:"data,omitempty"` // 场景下发 Message
}

var (
	redisPubSubManager *websocket.RedisPubSub
	once               sync.Once
)

func lazyInit() {
	once.Do(func() {
		pubSubChannels := make([]string, pubSubChannelCnt)
		for i := range pubSubChannels {
			pubSubChannels[i] = env.AppName() + ":pubSub:" + strconv.Itoa(i)
		}
		redisPubSubManager = websocket.NewRedisPubSub(pubSubChannels...)
	})
}

// 场景注册消息
// {
//  "scene": "test",
//  "sceneParams": {
//    "id": "11111111111111",
//  }
//}

// RedisPublish 存入样例
// {"code":0,"sid":"","scene":"test","id":"11111111111111","data":"xxxxxxxxxxx"}
func RedisPublish(ctx context.Context, msg Message) {
	lazyInit()
	msgBytes, _ := json.Marshal(msg)
	redisPubSubManager.Publish(ctx, cRedis.DefaultClient(), string(msgBytes))
}

func RedisSubscribe(ctx context.Context, m *websocket.ClientManager) (err error) {
	lazyInit()
	panic(redisPubSubManager.Subscribe(ctx, cRedis.DefaultClient(), func(fCtx context.Context, msg *redis.Message) {
		var pubSubMessage Message
		err = json.Unmarshal([]byte(msg.Payload), &pubSubMessage)
		// 针对 test 场景发送信息
		if pubSubMessage.Scene == SceneTest && len(pubSubMessage.ID) > 0 {
			m.Range(sceneTest(fCtx, pubSubMessage))
		}
	}))
}

func sceneTest(ctx context.Context, pubSubMessage Message) func(key, value any) bool {
	return func(key, value any) bool {
		if client, isOK := value.(*websocket.Client); isOK && SceneTest == client.State.Scene {
			if id, isExist := client.State.SceneParams["id"]; isExist {
				if idStr, isOK2 := id.(string); isOK2 && idStr == pubSubMessage.ID {
					pubSubMessage.State = client.State
					contentBytes, _ := json.Marshal(pubSubMessage)
					_ = client.WriteTextMessage(ctx, contentBytes)
				}
			}
		}
		return true
	}
}
