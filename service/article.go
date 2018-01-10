package service

import (
	. "os"
	"strconv"

	"github.com/JustinTulloss/firebase"
	"github.com/konojunya/riot-go-firebase/model"
)

var (
	endpoint string
	auth     string
	client   firebase.Client
)

func init() {
	endpoint = "https://riot-go-firebase.firebaseio.com"
	auth = Getenv("FIREBASE_AUTH_TOKEN")
	client = firebase.NewClient(endpoint+"/blog", auth, nil)
}

func articleAlloc() interface{} {
	return &model.Article{}
}

func GetArticleTitles() ([]string, error) {
	var titles []string

	for n := range client.Iterator(articleAlloc) {
		titles = append(titles, n.Value.(*model.Article).Title)
	}

	return titles, nil
}

func GetArticles() ([]model.Article, error) {
	var articles []model.Article

	for n := range client.Iterator(articleAlloc) {
		articles = append(articles, *n.Value.(*model.Article))
	}

	return articles, nil
}

func GetArticleById(id_str string) (*model.Article, error) {
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

	client.Push(article, nil)

	return nil
}

func UpdateArticle(id_str string, article *model.Article) error {
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
