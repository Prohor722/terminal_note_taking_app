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

func inputString(prompt string) string {
	for {
		fmt.Print(prompt)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Input cannot be empty.")
			continue
		}

		return input
	}
}

func inputInt(prompt string) int {
	for {
		fmt.Print(prompt)

		var id int
		_, err := fmt.Scanln(&id)
		if err != nil {
			fmt.Println("Invalid number. Try again.")
			clearInputBuffer()
			continue
		}

		return id
	}
}

func clearInputBuffer() {
	reader.ReadString('\n')
}