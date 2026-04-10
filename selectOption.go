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

// func selectOption(option int) bool {
//     switch option {
//     case 1:
//         caseCall(addNote)
//     case 2:
//         caseCall(viewNotes)
//     case 3:
//         caseCall(deleteNote)
//     case 4:
//         caseCall(editNote)
//     case 5:
//         fmt.Println("Exiting the Note Taking App...")
//         return false
//     default:
//         fmt.Println("Invalid option. Please try again.")
//     }
//     return true
// }

func caseCall(f func()) {
	f()
	pickOption()
}