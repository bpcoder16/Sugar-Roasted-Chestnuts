package bootstrap

import (
	"context"
	"github.com/bpcoder16/Chestnut/bootstrap"
	"github.com/bpcoder16/Chestnut/modules/appconfig"
)

func MustInit(ctx context.Context, config *appconfig.AppConfig) {
	bootstrap.MustInit(ctx, config)
}
