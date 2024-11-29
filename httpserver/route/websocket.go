package route

import (
	"Sugar-Roasted-Chestnuts/httpserver/websocket"
	"github.com/bpcoder16/Chestnut/appconfig/env"
	"github.com/bpcoder16/Chestnut/modules/ginwebsocket"
)

func WebSocket() *ginwebsocket.Router {
	webSocket := ginwebsocket.NewRouter(
		"/v1",
		env.RootPath()+"/conf/websocket.json",
	)
	webSocket.OnTextMessageController("test", &websocket.Test{})
	return webSocket
}
