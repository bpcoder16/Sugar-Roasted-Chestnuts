package route

import (
	"Sugar-Roasted-Chestnuts/httpserver/api"
	"github.com/bpcoder16/Chestnut/contrib/httphandler/gin"
	"net/http"
)

func Api() gin.Router {
	apiRouter := gin.NewDefaultRouter("/api")

	apiRouter.Any("/test", (&api.Test{}).Test)
	apiRouter.Match([]string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
	}, "/test2", (&api.Test{}).Test)

	testGroup := apiRouter.Group("/test")
	{
		testGroup.GET("/test", (&api.Test{}).Test)
	}

	return apiRouter
}

func ApiV2() gin.Router {
	apiRouter := gin.NewDefaultRouter("/api/v2")

	apiRouter.GET("/test", (&api.Test{}).Test)

	testGroup := apiRouter.Group("/test")
	{
		testGroup.GET("/test", (&api.Test{}).Test)
	}

	return apiRouter
}
