package logtest

import (
	"context"
	"github.com/bpcoder16/Chestnut/appconfig/env"
	"github.com/bpcoder16/Chestnut/logit"
)

func Test(ctx context.Context) {
	logit.Context(ctx).DebugW("env.LocalIPV4()", env.LocalIPV4())
	logit.Context(ctx).InfoW("env.LocalIPV4()", env.LocalIPV4())
	logit.Context(ctx).WarnW("env.LocalIPV4()", env.LocalIPV4())
	logit.Context(ctx).ErrorW("env.LocalIPV4()", env.LocalIPV4())
}
