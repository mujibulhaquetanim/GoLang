package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
	fmt.Printf("Todo %q added successfully at the index of %d\n\n", title, len(*todos)-1)
	fmt.Println("Your new Todos are: ")
	todos.print()
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos // dereferencing the pointer to get the value of the pointer. basically it copies the value of the pointer to a new variable
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	fmt.Printf("Todo: %q deleted successfully at the index of %d\n\n", t[index].Title, index)

	t = append(t[:index], t[index+1:]...) // removes the element from the slice and returns the new slice.
	*todos = t

	fmt.Println("Your new Todos are: ")
	todos.print()

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted
	fmt.Println("Your new state is", t[index].Completed, "for the index of", index)
	fmt.Printf("\nTodo %q updated successfully at the index of %d\n\n", t[index].Title, index)
	fmt.Println("Your new Todos are: ")
	todos.print()

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	fmt.Printf("Todo %q edited successfully at the index of %d\n\n", title, index)
	fmt.Println("Your new Todos are: ")
	todos.print()

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, todo := range *todos {
		completed := "❌"
		completedAt := ""

		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
