package clickhousetest

import (
	"Sugar-Roasted-Chestnuts/db"
	"context"
	"github.com/bpcoder16/Chestnut/clickhouse"
	"github.com/bpcoder16/Chestnut/core/utils"
	"github.com/bpcoder16/Chestnut/logit"
)

func Test(ctx context.Context) {
	clickhouse.MasterDB().WithContext(ctx).Create(&db.Test{
		Id:      utils.RandIntN(100000),
		Content: utils.RandIntStr(16),
		Status:  db.TestStatusDefault,
	})

	var test db.Test
	clickhouse.MasterDB().WithContext(ctx).First(&test)
	logit.Context(ctx).InfoW("clickhouse.MasterDB().First", test)
}
