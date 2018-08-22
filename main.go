package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weifuchuan/fuchuansia-server/controller"
)

func main() {
	server := gin.Default()
	//server.Static("/", "./webapp/")
	server.Static("/static", "./webapp/static")
	server.Static("/media", "./webapp/media")
	server.GET("/", func(c *gin.Context) {
		c.File("./webapp/index.html")
	})

	server.POST("/project/get", controller.GetProjects)

	server.POST("/project/add", controller.AddProject)

	server.POST("/auth", controller.Auth)

	server.POST("/media-upload", controller.UploadMedia)

	server.Run(":9000")
}
