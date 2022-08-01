package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() http.Handler {
	router := mux.NewRouter()
	// router.Get("/",)
	router.HandleFunc("/", RootEndPoint).Methods("GET")

	// Authors
	router.HandleFunc("/authors", AuthorRetriveAllEndpoint).Methods("GET")
	router.HandleFunc("/author", AuthorStoreEndpoint).Methods("POST")
	router.HandleFunc("/author/{id}", AuthorRetriveEndpoint).Methods("GET")
	router.HandleFunc("/author/{id}", AuthorDeleteEndpoint).Methods("DELETE")
	router.HandleFunc("/author/{id}", AuthorUpdateEndpoint).Methods("UPDATE")

	// Articles
	router.HandleFunc("/articles", ArticleRetriveAllEndpoint).Methods("GET")
	router.HandleFunc("/article", ArticleStoreEndpoint).Methods("POST")
	router.HandleFunc("/article/{id}", ArticleRetriveEndpoint).Methods("GET")
	router.HandleFunc("/article/{id}", ArticleDeleteEndpoint).Methods("DELETE")
	router.HandleFunc("/article/{id}", ArticleUpdateEndpoint).Methods("UPDATE")

	return router
}
