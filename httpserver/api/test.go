package api

import (
	"github.com/bpcoder16/Chestnut/logit"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Test struct{}

func (t *Test) Test(ctx *gin.Context) {
	logit.Context(ctx).InfoW("Test", "Test")
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "OK",
	})
}
