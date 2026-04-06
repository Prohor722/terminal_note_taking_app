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

var option = 0

func main() {
	fmt.Println("Welcome to the Note Taking App!")

	fmt.Println("Select a option:")
	fmt.Println("1. Add Note")
	fmt.Println("2. View Notes")
	fmt.Println("3. Delete Note")
	fmt.Println("4. Edit Note")
	fmt.Println("5. Exit")


}
