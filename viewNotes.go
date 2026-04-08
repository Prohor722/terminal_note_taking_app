package main

import "fmt"

func viewNotes() {
	if len(notes) == 0 {
		fmt.Println("No notes available.")
		return
	}
	fmt.Println("Your notes:")
	for i, n := range notes {
		fmt.Printf("No.%d \nID: %d\nTitle: %s\nBody: %s\n\n",i+1, n.id, n.title, n.body)
	}
}