package api

import (
	"github.com/bpcoder16/Chestnut/logit"
	"github.com/gin-gonic/gin"
)

type Test struct{}

func (t *Test) Test(ctx *gin.Context) {
	logit.Context(ctx).InfoW("Test", "Test")
}
