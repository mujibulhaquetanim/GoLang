package main

func main() {
	todos := Todos{}
	//type Todos passed as a generic type to NewStorage & fileName given as string
	storage := NewStorage[Todos]("todos.json")
	//passed a pointer of type Todos to Load
	storage.Load(&todos)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)


	storage.Save(todos)
}

// todos.print()
	// todos.add("Buy milk")
	// todos.add("Buy eggs")
	//toggles the todo at index 0 which will set it to completed
	// todos.toggle(0)
