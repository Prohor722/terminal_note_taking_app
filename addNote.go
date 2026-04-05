package main

import "fmt"

func addNote() {
	fmt.Println("Add new note..")
	fmt.Println("Enter note title:")
	var title string
	fmt.Scanln(&title)
	fmt.Println("Enter note body:")
	var body string
	fmt.Scanln(&body)
	note = struct {
		id  int
		title string
		body  string
	}{id: len(notes) + 1, title: title, body: body}
	notes = append(notes, struct {
		note struct {
			id  int
			title string
			body  string
		}
	}{note: note})
	fmt.Println("Note added successfully!")
}