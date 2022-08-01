package main

import "fmt"

func main() {
	a := "hello"
	v := fmt.Fprintf("%v world", a)
	v := fmt.Sprintf("%v world", a)
	fmt.Println(v)
	fmt.Println("Hello World, from FMT 👋")
}
