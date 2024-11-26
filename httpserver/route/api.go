package route

import (
	"Sugar-Roasted-Chestnuts/httpserver/api"
	"github.com/bpcoder16/Chestnut/contrib/httphandler/gin"
)

func Api() *gin.DefaultRouter {
	apiRouter := gin.NewDefaultRouter("/api")

	apiRouter.GET("/test", (&api.Test{}).Test)

	testGroup := apiRouter.Group("/test")
	{
		testGroup.GET("/test", (&api.Test{}).Test)
	}

	return apiRouter
}

func ApiV2() *gin.DefaultRouter {
	apiRouter := gin.NewDefaultRouter("/api/v2")

	apiRouter.GET("/test", (&api.Test{}).Test)

	testGroup := apiRouter.Group("/test")
	{
		testGroup.GET("/test", (&api.Test{}).Test)
	}

	return apiRouter
}
