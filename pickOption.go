package main

import "fmt"

func pickOption() {
	fmt.Println("Select a option:")
	fmt.Println("1. Add Note")
	fmt.Println("2. View Notes")
	fmt.Println("3. Delete Note")
	fmt.Println("4. Edit Note")
	fmt.Println("5. Exit")

	fmt.Print("Enter your choice: ")
	fmt.Scanln(&option)
	selectOption()
}