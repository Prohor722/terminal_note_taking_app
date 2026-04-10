package main

import "fmt"
var options = []string{
	"Add Note",
	"View Notes",
	"Delete Note",
	"Edit Note",
	"Exit",
}


func pickOption() int {
	var choice int

	fmt.Println("\nSelect an option:")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	for {
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanln(&choice)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if choice < 1 || choice > len(options) {
			fmt.Println("Please select a valid option.")
			continue
		}

		return choice
	}
}