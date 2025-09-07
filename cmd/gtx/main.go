package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var root = filepath.Join(os.Getenv("HOME"), ".config/gtx")

type cmd interface {
	name() string
	short() string
	long() string
}

var cmds = []cmd{ls, create, selectCommand, edit, help, rm, cat}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("not enough args")
		help.run(1, nil)
	}

	var err error
	switch os.Args[1] {
	case "help", "-h", "--help":
		help.run(0, os.Args[2:])
	case "ls":
		err = ls.run()
	case "select":
		err = selectCommand.run(os.Args[2:])
	case "create":
		err = create.run(os.Args[2:])
	case "edit":
		err = edit.run(os.Args[2:])
	case "rm", "del", "delete":
		err = rm.run(os.Args[2:])
	case "cat":
		err = cat.run(os.Args[2:])
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		help.run(1, nil)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, red.Sprint(err.Error()))
		os.Exit(1)
	}
}
