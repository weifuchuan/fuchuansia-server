package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/weifuchuan/fuchuansia-server/db"
	"github.com/weifuchuan/fuchuansia-server/kit"
)

func GetArticleBase(c *gin.Context) {
	req := make([]string, 0)
	if err := c.BindJSON(&req); err != nil {
		kit.Logger.Println(err)
		return
	}
	ids := make([]bson.ObjectId, 0)
	for i := range req {
		ids = append(ids, bson.ObjectIdHex(req[i]))
	}
	res := make([]H, 0)
	col := db.Articles()
	if err := col.Find(bson.M{
		"_id": H{
			"$in": ids,
		},
	}).Select(H{
		"content": 0,
	}).All(&res); err != nil {
		kit.Logger.Println(err)
		c.String(500, "error")
		return
	}
	for i := range res {
		res[i]["_id"] = res[i]["_id"].(bson.ObjectId).Hex()
		res[i]["articleClass"] = res[i]["articleClass"].(bson.ObjectId).Hex()
	}
	c.JSON(200, res)
}

func GetArticle(c *gin.Context) {
	req := make([]string, 0)
	if err := c.BindJSON(&req); err != nil {
		kit.Logger.Println(err)
		return
	}
	ids := make([]bson.ObjectId, 0)
	for i := range req {
		ids = append(ids, bson.ObjectIdHex(req[i]))
	}
	res := make([]H, 0)
	col := db.Articles()
	if err := col.Find(bson.M{
		"_id": H{
			"$in": ids,
		},
	}).All(&res); err != nil {
		kit.Logger.Println(err)
		c.String(500, "error")
		return
	}
	for i := range res {
		res[i]["_id"] = res[i]["_id"].(bson.ObjectId).Hex()
	}
	c.JSON(200, res)
}

func AddArticle(c *gin.Context) {
	req := &struct {
		Token        string `json:"token"`
		Title        string `json:"title"`
		CreateAt     int64  `json:"createAt"`
		Content      string `json:"content"`
		ArticleClass string `json:"articleClass"`
	}{}
	if err := c.BindJSON(req); err != nil {
		kit.Logger.Println(err)
		return
	}
	if req.Token != kit.Config.Token {
		c.String(500, "error")
		return
	}
	id := bson.NewObjectId()
	col := db.Articles()
	if err := col.Insert(H{"_id": id, "title": req.Title, "createAt": req.CreateAt, "content": req.Content, "articleClass": bson.ObjectIdHex(req.ArticleClass)}); err != nil {
		kit.Logger.Println(err)
		c.String(500, "error")
		return
	}
	acCol := db.ArticleClasses()
	if err := acCol.Update(H{
		"_id": bson.ObjectIdHex(req.ArticleClass),
	}, H{
		"$addToSet": H{"articles": id},
	}); err != nil {
		kit.Logger.Println(err)
		c.String(500, "error")
		return
	}
	c.JSON(200, H{"id": id.Hex()})
}

func UpdateArticle(c *gin.Context) {
	req := &struct {
		Token        string `json:"token"`
		Id           string `json:"_id"`
		Title        string `json:"title"`
		CreateAt     int64  `json:"createAt"`
		Content      string `json:"content"`
		ArticleClass string `json:"articleClass"`
	}{}
	if err := c.BindJSON(req); err != nil {
		kit.Logger.Println(err)
		return
	}
	if req.Token != kit.Config.Token {
		c.String(500, "error")
		return
	}
	col := db.Articles()
	old := H{}
	if err := col.Find(H{"_id": bson.ObjectIdHex(req.Id)}).One(&old); err != nil {
		kit.Logger.Println(err)
		c.String(500, "error")
		return
	}
	oldAC := old["articleClass"].(bson.ObjectId).Hex()
	if oldAC != req.ArticleClass {
		acCol := db.ArticleClasses()
		if err := acCol.Update(H{"_id": old["articleClass"].(bson.ObjectId)}, H{"$pull": H{"articles": H{"$eq": bson.ObjectIdHex(req.Id)}}}); err != nil {
			kit.Logger.Println(err)
			c.String(500, "error")
			return
		}
		if err := acCol.Update(H{"_id": bson.ObjectIdHex(req.ArticleClass)}, H{"$addToSet": H{"articles": bson.ObjectIdHex(req.Id)}}); err != nil {
			kit.Logger.Println(err)
			c.String(500, "error")
			return
		}
	}
	if err := col.Update(H{"_id": bson.ObjectIdHex(req.Id)}, H{"title": req.Title, "content": req.Content, "createAt": req.CreateAt, "articleClass": req.ArticleClass}); err != nil {
		kit.Logger.Println(err)
		c.String(500, "error")
		return
	}
	c.String(200, "")
}
