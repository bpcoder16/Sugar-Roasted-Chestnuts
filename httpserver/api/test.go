package api

import (
	"context"
	"github.com/bpcoder16/Chestnut/core/asynctask"
	"github.com/bpcoder16/Chestnut/logit"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Test struct{}

func (t *Test) Test(ctx *gin.Context) {
	logit.Context(ctx).InfoW("Test", "Test")

	asynctask.AddQueue(ctx, func(ctxQ context.Context) error {
		logit.Context(ctxQ).DebugW("Test", "Test")
		return nil
	}, "测试 ErrMsg")

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "OK",
	})
}
