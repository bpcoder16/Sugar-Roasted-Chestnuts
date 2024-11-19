package clickhousetest

import (
	SClickhouse "Sugar-Roasted-Chestnuts/clickhouse"
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

func MultipleTest(ctx context.Context) {
	SClickhouse.Test01MasterDB().WithContext(ctx).Create(&db.Test{
		Id:      utils.RandIntN(100000),
		Content: utils.RandIntStr(16),
		Status:  db.TestStatusDefault,
	})
	SClickhouse.Test02MasterDB().WithContext(ctx).Create(&db.Test{
		Id:      utils.RandIntN(100000),
		Content: utils.RandIntStr(16),
		Status:  db.TestStatusDefault,
	})

	var test01 db.Test
	SClickhouse.Test01MasterDB().WithContext(ctx).First(&test01)
	logit.Context(ctx).InfoW("SClickhouse.Test01MasterDB().First", test01)

	var test02 db.Test
	SClickhouse.Test02MasterDB().WithContext(ctx).First(&test02)
	logit.Context(ctx).InfoW("SClickhouse.Test02MasterDB().First", test02)
}
