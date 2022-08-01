package resources

type Article struct {
	Id      string `json:"id,omitempty"`
	Author  string `json:"author,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

func GetAllArticles() []Article {
	return getarticles()
}

func getarticles() []Article {
	articles := make([]Article, 0)

	article1 := Article{
		Id:      "1",
		Author:  "Mohamed",
		Title:   "Some blog Post",
		Content: "My first Blog post for an api",
	}

	article2 := Article{
		Id:      "2",
		Author:  "Ken Thompson",
		Title:   "Golang new PL",
		Content: "Structs in Golang",
	}

	articles = []Article{article1, article2}
	return articles
}
