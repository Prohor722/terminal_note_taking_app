package main

import "fmt"

var note struct {
	id   int
	title string
	body  string
}

var notes []struct {
	note struct {
		id  int
		title string
		body  string
	}
}

func main() {
	test()
	fmt.Println("Hello, World!")
}
