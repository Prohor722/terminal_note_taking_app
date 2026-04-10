package main

import "fmt"

type Note struct {
	id   int
	title string
	body  string
}

var notes []*Note

var option = 0

func main() {
	fmt.Println("Welcome to the Note Taking App!")
	pickOption()
	selectOption(option)
}
