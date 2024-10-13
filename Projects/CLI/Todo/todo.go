package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   string
	CompletedAt string
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now().String(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err:= errors.New("invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t:= *todos
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	t = append(t[:index], t[index+1:]...)
	*todos = t
	return nil
}