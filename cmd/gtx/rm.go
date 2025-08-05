package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type rmCmd struct{}

var rm = rmCmd{}

func (d rmCmd) name() string { return "rm" }

func (d rmCmd) short() string { return "Delete a context" }

func (d rmCmd) long() string {
	return `usage: rm REPO [CTX]

Delete an entire repo, or delete a context from
the repo specified
`
}

func (d rmCmd) run(args []string) error {
	switch len(args) {
	case 1:
		if ask("Are you sure you want to delete repo %s?", args[0]) {
			return os.RemoveAll(filepath.Join(root, args[0]))
		}
	case 2:
		return os.Remove(filepath.Join(root, args[0], args[1]+".yaml"))
	default:
		return fmt.Errorf("incorrect number of args to rm, must at least specify repo")
	}

	return nil
}

func ask(question string, args ...any) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", fmt.Sprintf(question, args...))

		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		response = strings.TrimSpace(strings.ToLower(response))

		switch response {
		case "y", "yes":
			return true
		default:
			return false
		}
	}
}
