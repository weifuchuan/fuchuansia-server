package controller

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/weifuchuan/fuchuansia-server/db"
	"github.com/weifuchuan/fuchuansia-server/kit"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type H = map[string]interface{}
type Any = interface{}

func GetProjects(c *gin.Context) {
	col := db.Projects()
	query := col.Find(nil)
	projects := make([]H, 0)
	err := query.All(&projects)
	if err != nil {
		log.Println(err)
		c.String(500, "error")
		return
	}
	for i := 0; i < len(projects); i++ {
		projects[i]["_id"] = projects[i]["_id"].(bson.ObjectId).Hex()
	}
	sort.Slice(projects, func(i, j int) bool {
		return projects[i]["order"].(float64) < projects[j]["order"].(float64)
	})
	c.JSON(200, H{"projects": projects})
}

func Auth(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		c.String(500, "error")
		return
	}
	var req struct {
		Token string `json:"token"`
	}
	json.Unmarshal(data, &req)
	if req.Token == kit.Config.Token {
		c.JSON(200, H{"result": "ok"})
	} else {
		kit.Logger.Println("Got token = " + req.Token)
		c.String(500, "error")
	}
}

func UploadMedia(c *gin.Context) {
	tkn := c.PostForm("token")
	if tkn != kit.Config.Token {
		c.String(http.StatusBadRequest, "error")
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "error")
		return
	}
	filename := Md5(file.Filename+time.Now().String()) + file.Filename[strings.LastIndex(file.Filename, "."):]
	err = c.SaveUploadedFile(file, filepath.Join(rootPath(), "webapp/media/"+filename))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "error")
		return
	}
	mediaUri := "/media/" + filename
	c.String(200, mediaUri)
}

func AddProject(c *gin.Context) {
	type Req struct {
		Token   string `json:"token"`
		Name    string `json:"name"`
		Icon    string `json:"icon"`
		Profile string `json:"profile"`
		Detail  string `json:"detail"`
		Order   string `json:"order"`
	}
	req := new(Req)
	if err := c.BindJSON(req); err != nil {
		log.Println(err)
		return
	}
	if req.Token != kit.Config.Token {
		c.String(500, "error")
	}
	projects := db.Projects()
	//_, err := projects.InsertOne(c, H{"name": req.Name, "icon": req.Icon, "profile": req.Profile, "detail": req.Detail})
	err := projects.Insert(bson.M{"name": req.Name, "icon": req.Icon, "profile": req.Profile, "detail": req.Detail, "order": req.Order})
	if err != nil {
		log.Println(err)
		c.String(500, "error")
		return
	}
	res := make(H)
	if err = projects.Find(bson.M{"name": req.Name}).One(&res); err != nil {
		log.Println(err)
		c.String(500, "error")
		return
	}
	c.JSON(200, H{"result": "ok", "id": res["_id"].(bson.ObjectId).Hex()})
}

func rootPath() string {
	p, _ := os.Getwd()
	return p
}

func Md5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func init() {
	log.SetFlags(log.Llongfile)
}
