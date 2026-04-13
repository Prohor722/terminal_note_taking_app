package main

import "fmt"

var reader = bufio.NewReader(os.Stdin)

func deleteNote() {
	fmt.Println("\n--- Delete Note ---")

	id := inputInt("Enter note ID to delete: ")

	index := findNoteIndexByID(id)
	if index == -1 {
		fmt.Println("❌ Note not found!")
		return
	}

	// Confirmation step (safer UX)
	if !confirmAction("Are you sure you want to delete this note? (y/n): ") {
		fmt.Println("❌ Deletion cancelled.")
		return
	}

	notes = append(notes[:index], notes[index+1:]...)
	fmt.Println("✅ Note deleted successfully!")
}

func findNoteIndexByID(id int) int {
	for i, n := range notes {
		if n.id == id {
			return i
		}
	}
	return -1
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

