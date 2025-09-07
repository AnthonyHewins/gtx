package main

import (
	"fmt"
	"path/filepath"

	"github.com/AnthonyHewins/gtx/pkg/gtx"
)

type selectCmd struct{}

func (s selectCmd) name() string {
	return "select"
}

func (s selectCmd) short() string {
	return "Select current active context"
}

func (s selectCmd) long() string {
	return `usage: select REPO CTX

Change the current active context in a particular repo
`
}

var selectCommand = selectCmd{}

func (s selectCmd) run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("must supply repo and env at minimum")
	}

	return gtx.SetCtx(filepath.Join(root, args[0]), args[1])
}
