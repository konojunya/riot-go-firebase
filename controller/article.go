package controller

import (
	"log"

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

func PostArticle(article *model.Article) {

	client := getClient()

	client.Push(article, nil)

}

func GetArticles() {

	client := getClient()

	for n := range client.Iterator(articleAlloc) {
		log.Println(n.Value.(*model.Article).Title)
	}

}

func UpdateArticle(article *model.Article) {

	client := getClient()

	for n := range client.Iterator(articleAlloc) {
		if n.Value.(*model.Article).Title == "テストタイトル" {
			client.Update(n.Key, article, nil)
		}
	}
}

func DeleteArticle(id uint) {

	client := getClient()

	for n := range client.Iterator(articleAlloc) {
		if n.Value.(*model.Article).ID == id {
			client.Remove(n.Key, nil)
		}
	}

}
