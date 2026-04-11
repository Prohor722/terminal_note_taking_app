package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addNote() {
	fmt.Println("\n--- Add New Note ---")

	title := inputWithValidation(
		"Enter note title: ",
		isTitleUnique,
		"Title already exists. Please enter a unique title.",
	)

	body := inputWithValidation(
		"Enter note body: ",
		isBodyUnique,
		"Body already exists. Please enter a unique body.",
	)

	note := &Note{
		id:    len(notes) + 1,
		title: title,
		body:  body,
	}

	notes = append(notes, note)
	fmt.Println("✅ Note added successfully!")
}

func inputWithValidation(prompt string, validate func(string) bool, errorMsg string) string {
	for {
		fmt.Print(prompt)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Input cannot be empty.")
			continue
		}

		if !validate(input) {
			fmt.Println(errorMsg)
			continue
		}

		return input
	}
}

func isTitleUnique(title string) bool {
	for _, note := range notes {
		if strings.EqualFold(note.title, title) {
			return false
		}
	}
	return true
}

func isBodyUnique(body string) bool {
	for _, note := range notes {
		if strings.EqualFold(note.body, body) {
			return false
		}
	}
	return true
}

// func addNote() {
// 	fmt.Println("Add new note..")
	
// 	title := addTitle()
// 	body := addBody()
	
// 	var note = Note{id: len(notes) + 1, title: title, body: body}
// 	notes = append(notes, &note)
// 	fmt.Println("Note added successfully!")
// }

// func addTitle() string{
// 	fmt.Println("Enter note title:")
// 	var title string
// 	fmt.Scanln(&title)

// 	if !validateTitle(title) {
// 		fmt.Println("Title already exists. Please enter a unique title.")
// 		return addTitle()
// 	}
// 	return title
// }

// func validateTitle(title string) bool {
// 	for _, note := range notes {
// 		if note.title == title {
// 			return false
// 		}
// 	}
// 	return true
// }

// func validateBody(body string) bool {
// 	for _, note := range notes {
// 		if note.body == body {
// 			return false
// 		}
// 	}
// 	return true
// }

// func addBody() string{
// 	fmt.Println("Enter note body:")
// 	var body string
// 	fmt.Scanln(&body)
// 	if !validateBody(body) {
// 		fmt.Println("Body already exists. Please enter a unique body.")
// 		return addBody()
// 	}
// 	return body
// }