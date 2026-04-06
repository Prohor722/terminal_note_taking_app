package main

import "fmt"

func selectOption() {
	switch option {
	case 1:
		addNote()
	case 2:
		viewNotes()
	case 3:
		deleteNote()
	case 4:
		editNote()
	case 5:
		fmt.Println("Exiting the Note Taking App...")
		return
	default:
		fmt.Println("Invalid option. Please try again.")
	}
}