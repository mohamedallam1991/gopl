package web

import (
	"encoding/json"
	"errors"
	"net/http"

	// "github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mohamedallam1991/gopl/apis/articles/resources"
	uuid "github.com/satori/go.uuid"
)

var articles []resources.Article = resources.GetAllArticles()

func ArticleRetriveAllEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	json.NewEncoder(response).Encode(articles)
}

func ArticleStoreEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var article resources.Article
	json.NewDecoder(request.Body).Decode(&article)

	article.Id = uuid.Must(uuid.NewV4(), errors.New("cant uuid must")).String()

	articles = append(articles, article)
	json.NewEncoder(response).Encode(article)
}

func ArticleRetriveEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)

	for _, article := range articles {
		if article.Id == params["id"] {
			json.NewEncoder(response).Encode(article)
			return
		}
	}

	json.NewEncoder(response).Encode(resources.Article{})
}

func ArticleUpdateEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	var changes resources.Article
	json.NewDecoder(request.Body).Decode(&changes)
	for index, article := range articles {
		if article.Id == params["id"] {
			if changes.Author != "" {
				article.Author = changes.Author
			}
			if changes.Content != "" {
				article.Content = changes.Content
			}
			if changes.Title != "" {
				article.Title = changes.Title
			}
			articles[index] = article
			json.NewEncoder(response).Encode(articles)
			return
		}
	}

	json.NewEncoder(response).Encode(resources.Article{})
}

func ArticleDeleteEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	for index, article := range articles {
		if article.Id == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			json.NewEncoder(response).Encode(articles)
			return
		}
	}
	json.NewEncoder(response).Encode(resources.Article{})
}
