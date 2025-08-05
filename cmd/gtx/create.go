package main

import (
	"fmt"

	"github.com/AnthonyHewins/gtx/internal/gtx"
)

type createCmd struct{}

var create = createCmd{}

func (c createCmd) name() string { return "create" }

func (c createCmd) short() string {
	return "Create a new context"
}

func (c createCmd) long() string {
	return `usage: create REPO [ENV]

Create a new context, and optionally also create a new environment
for that context
`
}

func (n *createCmd) run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("must give at least 1 arg for an env")
	}

	ctx, err := gtx.CreateCtx(root, args[0])
	if err != nil {
		return err
	}
	fmt.Printf("%s %s\n", bold.Sprint(args[0]), gray.Sprint("created"))

	if len(args) == 1 {
		return nil
	}

	if err = ctx.AddEnv(root, args[1]); err != nil {
		return err
	}
	fmt.Printf("%s %s %s\n", gray.Sprint("Env"), bold.Sprint(args[1]), gray.Sprint("created"))

	return nil
}
