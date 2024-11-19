package main

import (
	"Sugar-Roasted-Chestnuts/bootstrap"
	"Sugar-Roasted-Chestnuts/db"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bpcoder16/Chestnut/core/utils"
	"github.com/bpcoder16/Chestnut/logit"
	"github.com/bpcoder16/Chestnut/modules/appconfig"
	"github.com/bpcoder16/Chestnut/modules/appconfig/env"
	"github.com/bpcoder16/Chestnut/mysql"
)

func main() {
	config := appconfig.MustLoadAppConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootstrap.MustInit(ctx, config)

	fmt.Println(ctx)

	configBytes, _ := json.Marshal(config)
	fmt.Println(string(configBytes))
	fmt.Println(env.AppName())
	fmt.Println(env.RunMode())
	fmt.Println(env.RootPath())
	fmt.Println(env.LocalIPV4())

	logit.Context(ctx).DebugW("env.LocalIPV4()", env.LocalIPV4())
	logit.Context(ctx).InfoW("env.LocalIPV4()", env.LocalIPV4())
	logit.Context(ctx).WarnW("env.LocalIPV4()", env.LocalIPV4())
	logit.Context(ctx).ErrorW("env.LocalIPV4()", env.LocalIPV4())

	mysql.MasterDB().WithContext(ctx).Create(&db.Test{
		Content: utils.RandIntStr(16),
		Status:  &db.TestStatusDefault,
	})

	var test db.Test
	mysql.MasterDB().WithContext(ctx).First(&test)
	logit.Context(ctx).InfoW("mysql.MasterDB().First", test)

	var testSlave db.Test
	testSlave.Id = 2
	mysql.SlaveDB().WithContext(ctx).First(&testSlave)
	logit.Context(ctx).InfoW("mysql.SlaveDB().First", testSlave)

	testSlave.Content = utils.RandIntStr(32)
	mysql.MasterDB().WithContext(ctx).Save(&testSlave)
}
