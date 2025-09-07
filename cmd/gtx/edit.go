package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AnthonyHewins/gtx/pkg/gtx"
)

type editCmd struct{}

func (e editCmd) name() string { return "edit" }

func (e editCmd) short() string { return "Edit a context in a repo" }

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

	repo, err := gtx.CreateRepo(root, args[0])
	if err != nil {
		return err
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		return fmt.Errorf("no $EDITOR set; to edit this config, set this var to a text editor"+
			" or edit the file directly at %s", repo.Path())
	}

	run := exec.Command(editor, filepath.Join(root, args[0], args[1]+".yaml"))
	run.Stdin = os.Stdin
	run.Stdout = os.Stdout
	return run.Run()
}
