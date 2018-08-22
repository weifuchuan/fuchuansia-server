package db

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/weifuchuan/fuchuansia-server/kit"
	"log"
	"context"
)

var client *mongo.Client

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), kit.Config.Mongodb)
	if err != nil {
		log.Fatal(err)
	}
}

func Projects() *mongo.Collection {
	return client.Database("fuchuansia").Collection("projects")
}
