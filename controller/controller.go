package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weifuchuan/fuchuansia-server/db"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"path/filepath"
	"os"
	"time"
	"crypto/md5"
	"encoding/hex"
	"strings"
	"github.com/globalsign/mgo/bson"
	"github.com/weifuchuan/fuchuansia-server/kit"
)

const (
	token = "ecb268c2a71ecebe597dd1f9bc55244948ad2d9d2b99901d038155244948"
)

type H = map[string]interface{}

func GetProjects(c *gin.Context) {
	col := db.Projects()
	query := col.Find(nil)
	projects := make([]H, 0)
	err := query.All(&projects)
	if err != nil {
		kit.Logger.Println(err)
		c.String(500, "error")
		return
	}
	for i := 0; i < len(projects); i++ {
		projects[i]["_id"] = projects[i]["_id"].(bson.ObjectId).Hex()
	}
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
	if req.Token == token {
		c.JSON(200, H{"result": "ok"})
	} else {
		c.String(500, "error")
	}
}

func UploadMedia(c *gin.Context) {
	tkn := c.PostForm("token")
	if tkn != token {
		c.String(http.StatusBadRequest, "error")
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		kit.Logger.Println(err)
		c.String(http.StatusBadRequest, "error")
		return
	}
	filename := Md5(file.Filename+time.Now().String()) + file.Filename[strings.LastIndex(file.Filename, "."):]
	err = c.SaveUploadedFile(file, filepath.Join(rootPath(), "webapp/media/"+filename))
	if err != nil {
		kit.Logger.Println(err)
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
	}
	req := new(Req)
	if err := c.BindJSON(req); err != nil {
		kit.Logger.Println(err)
		return
	}
	if req.Token != token {
		c.String(500, "error")
	}
	projects := db.Projects()
	//_, err := projects.InsertOne(c, H{"name": req.Name, "icon": req.Icon, "profile": req.Profile, "detail": req.Detail})
	err := projects.Insert(bson.M{"name": req.Name, "icon": req.Icon, "profile": req.Profile, "detail": req.Detail})
	if err != nil {
		kit.Logger.Println(err)
		c.String(500, "error")
		return
	}
	c.JSON(200, H{"result": "ok"})
}

func rootPath() string {
	p, _ := os.Getwd()
	return p
}

func Md5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
