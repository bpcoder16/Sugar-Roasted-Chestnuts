package mongodbtest

import (
	SMongoDB "Sugar-Roasted-Chestnuts/mongodb"
	"context"
	"encoding/json"
	"github.com/bpcoder16/Chestnut/logit"
	"github.com/bpcoder16/Chestnut/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test(ctx context.Context) {
	filter := bson.D{
		{"title", bson.D{
			{"$exists", true},
			{"$ne", ""},
			{"$ne", nil},
		}},
	}
	cursor, err := mongodb.DefaultClient().Collection("xxx").Find(
		ctx,
		filter,
		options.Find().SetSort(bson.D{{"_id", 1}}).SetLimit(2),
	)
	defer func() {
		_ = cursor.Close(ctx)
	}()
	if err != nil {
		logit.Context(ctx).WarnW("mongodb.error", err)
		return
	}
	for cursor.Next(ctx) {
		var result bson.M // 使用 bson.M 来获取完整数据
		if errD := cursor.Decode(&result); errD != nil {
			logit.Context(ctx).WarnW("bson.M.Err", errD)
			continue
		}
		resultJ, _ := json.Marshal(result)
		logit.Context(ctx).InfoW("bson.M.Result", string(resultJ))
	}
}

func MultipleTest(ctx context.Context) {
	filter := bson.D{
		{"title", bson.D{
			{"$exists", true},
			{"$ne", ""},
			{"$ne", nil},
		}},
	}
	cursor, err := SMongoDB.Test01Client().Collection("xxx").Find(
		ctx,
		filter,
		options.Find().SetSort(bson.D{{"_id", 1}}).SetLimit(2),
	)
	defer func() {
		_ = cursor.Close(ctx)
	}()
	if err != nil {
		logit.Context(ctx).WarnW("SMongoDB.Test01Client().error", err)
		return
	}
	for cursor.Next(ctx) {
		var result bson.M // 使用 bson.M 来获取完整数据
		if errD := cursor.Decode(&result); errD != nil {
			logit.Context(ctx).WarnW("SMongoDB.Test01Client().bson.M.Err", errD)
			continue
		}
		resultJ, _ := json.Marshal(result)
		logit.Context(ctx).InfoW("SMongoDB.Test01Client().bson.M.Result", string(resultJ))
	}

	cursor, err = SMongoDB.Test02Client().Collection("xxx").Find(
		ctx,
		filter,
		options.Find().SetSort(bson.D{{"_id", 1}}).SetLimit(2),
	)
	defer func() {
		_ = cursor.Close(ctx)
	}()
	if err != nil {
		logit.Context(ctx).WarnW("SMongoDB.Test02Client().error", err)
		return
	}
	for cursor.Next(ctx) {
		var result bson.M // 使用 bson.M 来获取完整数据
		if errD := cursor.Decode(&result); errD != nil {
			logit.Context(ctx).WarnW("SMongoDB.Test02Client().bson.M.Err", errD)
			continue
		}
		resultJ, _ := json.Marshal(result)
		logit.Context(ctx).InfoW("SMongoDB.Test02Client().bson.M.Result", string(resultJ))
	}
}
