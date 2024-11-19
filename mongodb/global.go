package mongodb

import (
	"context"
	"github.com/bpcoder16/Chestnut/contrib/gomongodb"
	"github.com/bpcoder16/Chestnut/core/log"
	"go.mongodb.org/mongo-driver/mongo"
)

var defaultMongodbTest01Manager *gomongodb.Manager
var defaultMongodbTest02Manager *gomongodb.Manager

func SetTest01Manager(ctx context.Context, configPath string, logger *log.Helper) {
	defaultMongodbTest01Manager = gomongodb.NewManager(ctx, configPath, logger)
}

func Test01Client() *mongo.Database {
	return defaultMongodbTest01Manager.ClientDatabase()
}

func SetTest02Manager(ctx context.Context, configPath string, logger *log.Helper) {
	defaultMongodbTest01Manager = gomongodb.NewManager(ctx, configPath, logger)
}

func Test02Client() *mongo.Database {
	return defaultMongodbTest01Manager.ClientDatabase()
}
