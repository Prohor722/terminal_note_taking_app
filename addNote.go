package main

import "fmt"

func addNote() {
	fmt.Println("Add new note..")
	
	title := addTitle()
	body := addBody()
	
	var note = Note{id: len(notes) + 1, title: title, body: body}
	notes = append(notes, &note)
	fmt.Println("Note added successfully!")
}

func addTitle() string{
	fmt.Println("Enter note title:")
	var title string
	fmt.Scanln(&title)
	return title
}

func validateTitle(title string) bool {
	for _, note := range notes {
		if note.title == title {
			return false
		}
	}
	return true
}

func addBody() string{
	fmt.Println("Enter note body:")
	var body string
	fmt.Scanln(&body)
	return body
}