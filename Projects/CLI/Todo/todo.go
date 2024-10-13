package main

import "time"

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