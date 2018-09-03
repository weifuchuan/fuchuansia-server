package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/weifuchuan/fuchuansia-server/db"
	"github.com/weifuchuan/fuchuansia-server/kit"
	"sort"
)

func AddArticleClass(c *gin.Context) {
	type Req struct {
		Token   string `json:"token"`
		Name    string `json:"name"`
		Profile string `json:"profile"`
		Order   int    `json:"order"`
	}
	req := &Req{}
	if err := c.BindJSON(req); err != nil {
		kit.Logger.Println(err)
		return
	}
	if req.Token != kit.Config.Token {
		c.String(500, "error")
		return
	}
	col := db.ArticleClasses()
	if err := col.Insert(bson.M{"name": req.Name, "profile": req.Profile, "order": req.Order, "articles": []string{}}); err != nil {
		kit.Logger.Println(err)
		c.String(500, "error")
		return
	}
	cursor := col.Find(bson.M{"name": req.Name})
	res := H{}
	if err := cursor.One(&res); err != nil {
		kit.Logger.Println(err)
		c.String(500, "error")
		return
	}
	c.JSON(200, H{"id": res["_id"].(bson.ObjectId).Hex()})
}

func GetArticleClass(c *gin.Context) {
	col := db.ArticleClasses()
	cursor := col.Find(nil)
	res := make([]H, 0)
	if err := cursor.All(&res); err != nil {
		kit.Logger.Println(err)
		c.String(500, "error")
		return
	}
	for i := range res {
		res[i]["_id"] = res[i]["_id"].(bson.ObjectId).Hex()
		ids := make([]string, 0)
		for _, id := range res[i]["articles"].([]Any) {
			ids = append(ids, id.(bson.ObjectId).Hex())
		}
		res[i]["articles"] = ids
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i]["order"].(int) < res[j]["order"].(int)
	})
	c.JSON(200, H{"articleClasses": res})
}
