package route

import (
	"github.com/bpcoder16/Chestnut/contrib/httphandler/gin"
	"github.com/bpcoder16/Chestnut/modules/ginwebsocket"
)

func WebSocket() gin.Router {
	webSocket := ginwebsocket.NewRouter(
		"/v1",
		"/conf/websocket.json",
	)
	return webSocket
}
