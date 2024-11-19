package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bpcoder16/Chestnut/modules/appconfig"
	"github.com/bpcoder16/Chestnut/modules/appconfig/env"
)

func main() {
	config := appconfig.MustLoadAppConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println(ctx)

	configBytes, _ := json.Marshal(config)
	fmt.Println(string(configBytes))
	fmt.Println(env.AppName())
	fmt.Println(env.RunMode())
	fmt.Println(env.RootPath())
	fmt.Println(env.LocalIPV4())
}
