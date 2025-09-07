package main

import (
	"fmt"

	"github.com/AnthonyHewins/gtx/pkg/gtx"
)

var cat = catCmd{}

type catCmd struct{}

func (c catCmd) name() string { return "cat" }

func (c catCmd) short() string { return "Cat out the contents of a context" }

func (c catCmd) long() string {
	return `Cats out the contents of a ctx`
}

func (c catCmd) run(args []string) error {
	n := len(args)
	switch n {
	case 0:
		return fmt.Errorf("missing args")
	case 1, 2:
	default:
		return fmt.Errorf("too many args: %s", args)
	}

	repo, err := gtx.ReadRepo(gtx.DefaultRoot, args[0])
	if err != nil {
		return err
	}

	var buf []byte
	if n == 1 {
		buf, err = repo.ReadCurrent()
	} else {
		buf, err = repo.Read(args[1])
	}

	if err != nil {
		return err
	}

	fmt.Print(string(buf))
	return nil
}
