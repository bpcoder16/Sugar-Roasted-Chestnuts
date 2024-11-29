package websocket

import (
	"context"
	"github.com/bpcoder16/Chestnut/contrib/websocket"
)

type Test struct {
	websocket.BaseTextMessageController
}

func (t *Test) Process(ctx context.Context) (err error) {
	return t.Client.WriteTextMessage(ctx, []byte("testtest"))
}
