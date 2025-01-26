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
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new task by giving title")
	flag.IntVar(&cf.Del, "del", -1, "Delete a task by index")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a task by '[index]:[new task]'")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Edit a task by index and input changes")
	flag.BoolVar(&cf.List, "list", false, "Toggle a task by index")

	flag.Parse()

	return &cf

}

func (cf *CmdFlags) Execute(todo *Task) {
	switch {
	case cf.List:
		todo.print()
	case cf.Add != "":
		todo.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid fromat. Pls use '[index]:[new task]'")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error, Invalid index for edit")
			os.Exit(1)
		}
		todo.change(index, parts[1])
	case cf.Toggle != -1:
		todo.toggle(cf.Toggle)
	case cf.Del != -1:
		todo.delete(cf.Del)
	default:
		fmt.Println("Invalid command")
	}
}
