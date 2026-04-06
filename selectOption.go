package main

import "fmt"

func selectOption() {
	switch option {
	case 1:
		addNote()
		pickOption()
	case 2:
		viewNotes()
		pickOption()
	case 3:
		deleteNote()
		pickOption()
	case 4:
		editNote()
		pickOption()
	case 5:
		fmt.Println("Exiting the Note Taking App...")
		return
	default:
		fmt.Println("Invalid option. Please try again.")
		pickOption()
	}
}