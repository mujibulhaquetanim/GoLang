package main

import "fmt"

func main() {
	// Greeting message
	fmt.Printf("\n ✨ Welcome to Todo CLI Application! ✨\n\n")
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json") // created a new instance of Storage wtih fileName & generic type
	storage.Load(&todos)                       // passed a pointer of type Todos to the Load function

	cmdFlags := NewCmdFlags() // created a new instance of CmdFlags
	cmdFlags.Execute(&todos)  // passed a pointer of type Todos to the Execute function
	storage.Save(todos)       // passed a value of type Todos to the Save function

	fmt.Println("\nFor detailed information/help, visit: https://github.com/mujibulhaquetanim/GoLang/tree/main/Projects/CLI/Todo")
	fmt.Print("\nPress **Enter** if you want to exit...")

	// Prevent the terminal from closing immediately
	fmt.Scanln()
}
