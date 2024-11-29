package websocket

import (
	"context"
	"github.com/bpcoder16/Chestnut/contrib/websocket"
	"github.com/bpcoder16/Chestnut/logit"
)

type Test struct {
	websocket.BaseTextMessageController
}

func (t *Test) Process(ctx context.Context) (err error) {
	logit.Context(ctx).DebugW("test", "test")
	return t.Client.WriteTextMessage(ctx, []byte("testtest"))
}
