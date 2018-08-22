package db

import (
	"github.com/weifuchuan/fuchuansia-server/kit"
	"github.com/globalsign/mgo"
)

var session *mgo.Session

func init() {
	var err error
	session, err = mgo.Dial(kit.Config.Mongodb)
	if err != nil {
		kit.Logger.Fatal(err)
	}
}

func Projects() *mgo.Collection {
	return session.DB("fuchuansia").C("projects")
}
