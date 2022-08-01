package main

import (
	"net/http"

	"github.com/mohamedallam1991/gopl/apis/articles/web"
)

// // slices of author
// var authors []resources.Author

// // slices of article
// var articles []resources.Article

const PORT = ":8765"

func main() {
	// router := mux.NewRouter()
	// router.HandleFunc("")
	http.ListenAndServe(PORT, web.Router())
}
