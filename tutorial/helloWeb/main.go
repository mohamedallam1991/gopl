package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8081"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		switch url {
		case "/":
			home(w, r)
		case "/about":
			about(w, r)
		case "/contact":
			contact(w, r)
		}
	})

	http.ListenAndServe(port, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from about me")
}

func contact(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		name := r.FormValue("name")
		email := r.FormValue("email")
		fmt.Fprintf(w, "Hello %s, your email: %s\n", name, email)
		fmt.Printf("Hello %s, your email: %s\n", name, email)
	default:
		fmt.Fprintf(w, "from contact page")
	}
}
