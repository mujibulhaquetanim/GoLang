package main

import "flag"

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
