package main

import "fmt"
var options = []string{
	"Add Note",
	"View Notes",
	"Delete Note",
	"Edit Note",
	"Exit",
}

func pickOption() {
	fmt.Println("Select a option:")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	fmt.Print("Enter your choice: ")
	fmt.Scanln(&option)
	selectOption()
}