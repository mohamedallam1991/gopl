package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/contact/*", contact)
	http.ListenAndServe(":8081", nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Path
	strings.Split(a, "/")
	fmt.Println(a)
	switch r.Method {
	case "GET":
		fetch(w, a)
	case "POST":
		a := r.FormValue("user")
		fmt.Println(a)
		fmt.Fprintf(w, "from POST")
	}
}

func fetch(w http.ResponseWriter, a string) {
	fmt.Fprintf(w, "from get, you are user %s", a)
}

func home(w http.ResponseWriter, r *http.Request) {
	_ = r
	fmt.Fprintf(w, "hello")
}
