package gtx

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AnthonyHewins/gtx/internal/dir"
	"github.com/goccy/go-yaml"
)

var (
	DefaultRoot = filepath.Join(os.Getenv("HOME"), ".config", "gtx")

	ErrNoCurrentCtx = errors.New("no current context")
)

// Calls ReadIntoFrom using DefaultRoot
func ReadInto(repo string, x any) error {
	return ReadIntoFrom(DefaultRoot, repo, x)
}

// Read the current config located at root's repo
// into x
func ReadIntoFrom(root, repo string, x any) error {
	buf, err := readCurrent(root, repo)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(buf, x)
}

func ReadIntoEnvFrom(root, repo string) error {
	return nil
}

func ReadIntoEnv(repo string) {
	ReadIntoEnvFrom(
		filepath.Join(os.Getenv("HOME"), ".config", "gtx"),
		repo,
	)
}

func readCurrent(path, repo string) ([]byte, error) {
	c, err := dir.ReadCtx(path, repo)
	if err != nil {
		return nil, err
	}

	if c.Current == "" {
		return nil, fmt.Errorf(
			"ctx %s in %s: %w. To add: echo '%s' > %s",
			repo,
			path,
			ErrNoCurrentCtx,
			repo,
			filepath.Join(path, repo, "current"),
		)
	}

	return c.Read()
}
