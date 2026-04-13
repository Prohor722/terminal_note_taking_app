package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader2 = bufio.NewReader(os.Stdin)

func deleteNote() {
	fmt.Println("\n--- Delete Note ---")

	id := inputInt2("Enter note ID to delete: ")

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

func inputInt2(prompt string) int {
	for {
		fmt.Print(prompt)

		var id int
		_, err := fmt.Scanln(&id)
		if err != nil {
			fmt.Println("Invalid number. Try again.")
			clearInputBuffer2()
			continue
		}
		return id
	}
}

func confirmAction(prompt string) bool {
	for {
		fmt.Print(prompt)

		input, _ := reader2.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "y" || input == "yes" {
			return true
		}
		if input == "n" || input == "no" {
			return false
		}

		fmt.Println("Please enter 'y' or 'n'.")
	}
}

func clearInputBuffer2() {
	reader2.ReadString('\n')
}