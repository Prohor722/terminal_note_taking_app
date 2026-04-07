package main

import "fmt"

func selectOption() {
	switch option {
	case 1:
		caseCall(addNote)
	case 2:
		caseCall(viewNotes)
	case 3:
		caseCall(deleteNote)
	case 4:
		caseCall(editNote)
	case 5:
		fmt.Println("Exiting the Note Taking App...")
		return
	default:
		fmt.Println("Invalid option. Please try again.")
	}
}

func caseCall(f func()) {
	f()
	pickOption()
}