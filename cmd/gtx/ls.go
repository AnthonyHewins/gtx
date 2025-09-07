package main

import (
	"fmt"

	"github.com/AnthonyHewins/gtx/pkg/gtx"
)

type lsCmd struct{}

var ls = lsCmd{}

func (l lsCmd) name() string { return "ls" }

func (l lsCmd) short() string { return "List information about contexts" }

func (l lsCmd) long() string {
	return `usage: ls

Lists information about contexts from config root
`
}

func (l lsCmd) run() error {
	t, err := gtx.NewTree(root)
	if err != nil {
		return err
	}

	bold.Printf("Config at %s\n\n", root)
	for i, c := range t.Ctxs {
		bold.Println(c.Repo)

		if len(c.Envs) == 0 {
			gray.Printf("No contexts\n\n")
			continue
		}

		for j, env := range c.Envs {
			leaf := "├──"
			if j == len(c.Envs)-1 {
				leaf = "└──"
			}

			envStr := gray.Sprint(env)
			if c.Current == env {
				envStr += green.Sprint(" *")
			}

			fmt.Printf("%s %s\n", leaf, envStr)
		}

		if i != len(t.Ctxs)-1 {
			fmt.Println()
		}
	}

	return nil
}
