package main

import (
	"fmt"

	"github.com/mohamedallam1991/gopl/apis/articles/resources"
)

// slices of author
var authors []resources.Author

// slices of article
var articles []resources.Article

const PORT = ":8765"

func main() {
	fmt.Println("starting the app")
	fmt.Printf("%T %v", getauthors(), getauthors())

}

func getarticles() []resources.Author {
	author1 := resources.Author{
		Id:        "author-1",
		Firstname: "Christiano",
		Lastname:  "Ronaldo",
		Username:  "nraboy",
		Password:  "pass",
	}

	author2 := resources.Author{
		Id:        "author-2",
		Firstname: "Maria",
		Lastname:  "Rabboy",
		Username:  "mraboy",
		Password:  "abc123",
	}
	authors = []resources.Author{author1, author2}
	return authors

}
func getauthors() []resources.Article {
	a := resources.Article{
		Id:      "article-1",
		Author:  "author-1",
		Title:   "blog post 1",
		Content: "This is our first blog post",
	}
	articles = []resources.Article{a, a}
	return articles
}
