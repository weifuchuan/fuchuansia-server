package db

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/weifuchuan/fuchuansia-server/kit"
	"log"
	"github.com/globalsign/mgo"
)

var client *mongo.Client
var session *mgo.Session

func init() {
	var err error
	session, err = mgo.Dial(kit.Config.Mongodb)
	if err != nil {
		log.Fatal(err)
	}
}

func Projects() *mgo.Collection {
	return session.DB("fuchuansia").C("projects")
}
