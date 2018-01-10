package controller

import (
	"log"
	"testing"

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

func TestPostArticle(t *testing.T) {
	client := getClient()
	client.Push(&model.Article{
		Title: "タイトル",
		Text:  "テキスト",
	}, nil)
}

func TestGetArticles(t *testing.T) {
	client := getClient()

	for n := range client.Iterator(articleAlloc) {
		log.Println(n.Value.(*model.Article).Title)
	}
}

func TestUpdateArticle(t *testing.T) {
	client := getClient()

	for n := range client.Iterator(articleAlloc) {
		if n.Value.(*model.Article).Title == "テストタイトル" {
			client.Update(n.Key, &model.Article{
				Title: "タイトル2",
				Text:  "テキスト",
			}, nil)
		}
	}
}

func TestDeleteArticle(t *testing.T) {
	var id uint
	id = 0

	client := getClient()

	for n := range client.Iterator(articleAlloc) {
		if n.Value.(*model.Article).ID == id {
			client.Remove(n.Key, nil)
		}
	}
}
