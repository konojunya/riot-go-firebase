package service

import (
	"strconv"

	"github.com/JustinTulloss/firebase"
	"github.com/konojunya/riot-go-firebase/model"
)

const (
	endpoint = "https://riot-go-firebase.firebaseio.com"
	auth     = "XFRU00YuHk9YOTh8dqgkAoMn1BDJsmGLhQMjRwFK"
)

func getClient() firebase.Client {
	return firebase.NewClient(endpoint+"/foo", auth, nil)
}

func articleAlloc() interface{} {
	return &model.Article{}
}

func GetArticles() ([]model.Article, error) {
	client := getClient()
	var articles []model.Article

	for n := range client.Iterator(articleAlloc) {
		articles = append(articles, *n.Value.(*model.Article))
	}

	return articles, nil
}

func GetArticleById(id_str string) (*model.Article, error) {
	client := getClient()
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return nil, err
	}
	var article *model.Article

	for n := range client.Iterator(articleAlloc) {
		if n.Value.(*model.Article).ID == uint(id) {
			article = n.Value.(*model.Article)
		}
	}

	return article, nil
}

func PostArticle(article *model.Article) error {
	client := getClient()

	client.Push(article, nil)

	return nil
}

func UpdateArticle(id_str string, article *model.Article) error {
	client := getClient()
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return err
	}

	for n := range client.Iterator(articleAlloc) {
		if n.Value.(*model.Article).ID == uint(id) {
			client.Update(n.Key, article, nil)
		}
	}

	return nil
}

func DeleteArticle(id_str string) error {
	client := getClient()
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return err
	}

	for n := range client.Iterator(articleAlloc) {
		if n.Value.(*model.Article).ID == uint(id) {
			client.Remove(n.Key, nil)
		}
	}

	return nil
}
