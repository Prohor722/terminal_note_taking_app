package main

import "fmt"


func editNote() {
	fmt.Println("\n--- Edit Note ---")

	id := inputInt("Enter note ID to edit: ")

	note := findNoteByID(id)
	if note == nil {
		fmt.Println("❌ Note not found!")
		return
	}

	fmt.Println("Editing note:", note.title)

	title := inputString("Enter new title: ")
	body := inputString("Enter new body: ")

	note.title = title
	note.body = body

	fmt.Println("✅ Note updated successfully!")
}

func findNoteByID(id int) *Note {
	for _, n := range notes {
		if n.id == id {
			return n
		}
	}
	return nil
}
