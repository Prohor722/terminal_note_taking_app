package main

import "fmt"

func addTitle() string {
	var title string
	fmt.Print("Enter title: ")
	fmt.Scanln(&title)
	return title
}

func editNote() {
	fmt.Println("Edit note..")
	fmt.Println("Enter note id to edit:")
	var id int
	fmt.Scanln(&id)
	for i, n := range notes {
		if n.id == id {
			title := addTitle()
			body := addBody()

			notes[i].title = title
			notes[i].body = body
			fmt.Println("Note updated successfully!")
			return
		}
	}
	fmt.Println("Note not found!")
}
