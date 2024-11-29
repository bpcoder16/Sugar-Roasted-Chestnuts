package route

import (
	"Sugar-Roasted-Chestnuts/httpserver/websocket"
	"github.com/bpcoder16/Chestnut/appconfig/env"
	"github.com/bpcoder16/Chestnut/contrib/httphandler/gin"
	"github.com/bpcoder16/Chestnut/modules/ginwebsocket"
)

func WebSocket() gin.Router {
	webSocket := ginwebsocket.NewRouter(
		"/v1",
		env.RootPath()+"/conf/websocket.json",
	)
	webSocket.OnTextMessageController("test", &websocket.Test{})
	return webSocket
}
