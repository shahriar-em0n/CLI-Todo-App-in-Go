package main

import "flag"

type cmdFlags struct {
	Add    string
	Delete int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")

	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")

	flag.IntVar(&cf.Delete, "edit", -1, "Specify a todo by index to delete")

	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")

	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}
