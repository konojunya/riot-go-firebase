package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/konojunya/riot-go-firebase/model"
	"github.com/konojunya/riot-go-firebase/service"
)

func GetArticles(c *gin.Context) {
	articleTitles, err := service.GetArticleTitles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, articleTitles)
}

func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	article, err := service.GetArticleById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, article)
}

func PostArticle(c *gin.Context) {
	title := c.PostForm("title")
	text := c.PostForm("text")

	err := service.PostArticle(&model.Article{
		Title: title,
		Text:  text,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	text := c.PostForm("text")

	err := service.UpdateArticle(id, &model.Article{
		Title: title,
		Text:  text,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	err := service.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
