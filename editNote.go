package main

import "fmt"

func editNote() {
	fmt.Println("Edit note..")
	fmt.Println("Enter note id to edit:")
	var id int
	fmt.Scanln(&id)
	for i, n := range notes {
		if n.note.id == id {
			title := addTitle()
			body := addBody()
			
			notes[i].note.title = title
			notes[i].note.body = body
			fmt.Println("Note updated successfully!")
			return
		}
	}
	fmt.Println("Note not found!")
}