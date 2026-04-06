package main

import "fmt"

func deleteNote(){
	fmt.Println("Delete a note..")
	fmt.Println("Enter note id to delete:")
	var id int
	fmt.Scanln(&id)
	for i, n := range notes {
		if n.note.id == id {
			notes = append(notes[:i], notes[i+1:]...)
			fmt.Println("Note deleted successfully!")
			return
		}
	}
	fmt.Println("Note not found!")
}