package storage

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/wumuwumu/trace-cloud/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var MongodbClient *mongo.Client

func Init(config config.Config){
	var err error
	MongodbClient, err = mongo.NewClient(options.Client().ApplyURI(config.MongoDB.URI))
	if err != nil{
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err = MongodbClient.Connect(ctx);err != nil{
		panic(err)
	}

	for{
		ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
		err = MongodbClient.Ping(ctx, readpref.Primary())
		if err == nil{
			break
		}
	}
	logrus.Info("MongoDB连接成功")

}

func Database() *mongo.Database{
	return MongodbClient.Database("TraceCloud")
}
