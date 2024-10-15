package main

func main() {
	todos := Todos{}

	todos.add("Buy milk")
	todos.add("Buy eggs")

	//toggles the todo at index 0 which will set it to completed
	todos.toggle(0)
	todos.print()
}
