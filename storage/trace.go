package storage

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Trace struct {
	TraceId string `bson:"traceId",json:"traceId"`
}

func Collection() *mongo.Collection{
	return Database().Collection("Trace")
}

func CreateTrace(trace Trace)  (interface{},error){
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := Collection().InsertOne(ctx, trace)
	if err != nil{
		return nil,err
	}
	return res.InsertedID,nil
}

func GetTrace(traceId string)(*Trace,error){
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var trace Trace
	filter := bson.D{{"traceId", traceId}}
	err := Collection().FindOne(ctx, filter).Decode(&trace)
	if err != nil {
		logrus.Fatal(err)
		return nil,err
	}
	return &trace,nil


}