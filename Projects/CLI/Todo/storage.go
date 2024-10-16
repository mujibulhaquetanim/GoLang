package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

// constructor for Storage
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{
		FileName: fileName,
	}
}

func (s *Storage[T]) Save(data T) error {
	//convert data to json, "\t" for indentation & "" used for prefix for each line.
	fileData, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return err
	}
	//0644 = read and write permission for owner only & read permission for others
	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}
	//convert json to data & store in data as data is a pointer of T
	//The json.Unmarshal function specifically expects a []byte as its first argument.
	//Using []byte allows for more efficient memory handling in some cases, especially when dealing with large amounts of data.
	return json.Unmarshal(fileData, data)
}
