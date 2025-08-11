package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AnthonyHewins/gtx/internal/dir"
)

type editCmd struct{}

func (e editCmd) name() string {
	return "edit"
}

func (e editCmd) short() string {
	return "Edit a context in a repo"
}

func (e editCmd) long() string {
	return fmt.Sprintf(`usage: edit REPO CTX

Edit a particular context in a repo.
If the repo/context doesn't exist, all will be created.
Looks in %s`,
		bold.Sprint(root),
	)
}

var edit = editCmd{}

func (e editCmd) run(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("not enough args to edit")
	}

	_, err := dir.CreateCtx(root, args[0])
	if err != nil {
		return err
	}

	run := exec.Command(os.Getenv("EDITOR"), filepath.Join(root, args[0], args[1]+".yaml"))
	run.Stdin = os.Stdin
	run.Stdout = os.Stdout
	return run.Run()
}
