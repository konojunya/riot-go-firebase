package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/riot-go-firebase/controller"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/images", "./public/images")

	r.LoadHTMLGlob("view/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	api := r.Group("/api")
	api.GET("/articles", controller.GetArticles)
	api.POST("/articles", controller.PostArticle)
	api.PUT("/articles/:id", controller.UpdateArticle)
	api.DELETE("/articles/:id", controller.DeleteArticle)

	r.Run(":3000")
}
