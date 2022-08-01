package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "https://google.com/articles/some-article-new-on-golang/more/splitting/just/to/make/sure"

	splitted := strings.Split(str, "/")[4]
	article := strings.Split(splitted, "-")
	fmt.Println(len(str))
	fmt.Println(len(splitted))
	fmt.Println(splitted[1])
	fmt.Println(splitted)
	fmt.Println(article)
	fmt.Println(str)
}
