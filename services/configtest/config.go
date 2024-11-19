package configtest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bpcoder16/Chestnut/modules/appconfig"
	"github.com/bpcoder16/Chestnut/modules/appconfig/env"
)

func Test(ctx context.Context, config *appconfig.AppConfig) {
	fmt.Println(ctx)

	configBytes, _ := json.Marshal(config)
	fmt.Println(string(configBytes))
	fmt.Println(env.AppName())
	fmt.Println(env.RunMode())
	fmt.Println(env.RootPath())
	fmt.Println(env.LocalIPV4())
}
