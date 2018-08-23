package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weifuchuan/fuchuansia-server/controller"
	"github.com/weifuchuan/fuchuansia-server/kit"
	"fmt"
	"os"
)

func main() {
	os.Mkdir("./webapp/media", os.ModeDir)

	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.File("./webapp/index.html")
	})

	server.POST("/project/get", controller.GetProjects)

	server.POST("/project/add", controller.AddProject)

	server.POST("/auth", controller.Auth)

	server.POST("/media-upload", controller.UploadMedia)

	server.Run(":" + fmt.Sprint(kit.Config.Port))
}
