package controller

import (
	"testing"

	"github.com/konojunya/riot-go-firebase/model"
)

func TestPostArticle(t *testing.T) {
	PostArticle(&model.Article{
		Title: "タイトル",
		Text:  "テキスト",
	})
}

func TestGetArticles(t *testing.T) {
	GetArticles()
}

func TestUpdateArticle(t *testing.T) {
	UpdateArticle(&model.Article{
		Title: "タイトル2",
		Text:  "テキスト",
	})
}

func TestDeleteArticle(t *testing.T) {
	var id uint
	id = 0

	DeleteArticle(id)
}
