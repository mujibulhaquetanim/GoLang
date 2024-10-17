package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Toggle int
	Delete int
	List   bool
	Edit   string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new Todo specify title")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit a new Todo by index & specify a new title. id: new_title")
	flag.IntVar(&cf.Delete, "Delete", -1, "Specify a todo by index to delete")
	flag.IntVar(&cf.Toggle, "Toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "List", false, "List All todos")

	//flag.Parse() is used to parse command line arguments passed to the program from terminal and store them in a struct variable.
	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()

	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, Invalid for Editing. Please use id: new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit")
			os.Exit(1)
		}
		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Delete != -1:
		todos.delete(cf.Delete)

	default:
		fmt.Println("Invalid Command")
	}
}
