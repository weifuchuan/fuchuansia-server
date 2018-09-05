package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weifuchuan/fuchuansia-server/controller"
	"github.com/weifuchuan/fuchuansia-server/kit"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	//router.Static("/static", "./webapp/static")

	router.GET("/", func(c *gin.Context) {
		c.File("./webapp/index.html")
	})

	group := router.Group("/project")
	{
		group.POST("/get", controller.GetProjects)
		group.POST("/add", controller.AddProject)
	}

	router.POST("/auth", controller.Auth)

	router.POST("/media-upload", controller.UploadMedia)

	group = router.Group("/article-class")
	{
		group.POST("/get", controller.GetArticleClass)
		group.POST("/add", controller.AddArticleClass)
	}

	group = router.Group("/article")
	{
		group.POST("/get/base", controller.GetArticleBase)
		group.POST("/get", controller.GetArticle)
		group.POST("/add", controller.AddArticle)
		group.POST("/update", controller.UpdateArticle)
	}

	router.Run(":" + fmt.Sprint(kit.Config.Port))
}
