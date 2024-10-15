package main

type Storage[T any] struct {
	FileName string
}

// constructor for Storage
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{
		FileName: fileName,
	}
}
