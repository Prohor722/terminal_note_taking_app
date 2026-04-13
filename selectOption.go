package main

import "fmt"

const (
    OptionAdd    = 1
    OptionView   = 2
    OptionDelete = 3
    OptionEdit   = 4
    OptionExit   = 5
)

var actions = map[int]func(){
    OptionAdd:    addNote,
    OptionView:   viewNotes,
    OptionDelete: deleteNote,
    OptionEdit:   editNote,
}

func selectOption(option int) bool {
    if option == OptionExit {
        fmt.Println("Exiting the Note Taking App...")
        return false
    }
    if action, ok := actions[option]; ok {
        caseCall(action)
    } else {
        fmt.Println("Invalid option. Please try again.")
    }
    return true
}

func caseCall(f func()) {
	f()
	if pickOption()>0 || pickOption()<6 {
        selectOption(option)
    }
}