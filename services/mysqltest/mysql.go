package mysqltest

import (
	"Sugar-Roasted-Chestnuts/db"
	"context"
	"github.com/bpcoder16/Chestnut/core/utils"
	"github.com/bpcoder16/Chestnut/logit"
	"github.com/bpcoder16/Chestnut/mysql"
)

func Test(ctx context.Context) {
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
