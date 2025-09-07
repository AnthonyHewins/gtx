package gtx

import (
	"errors"
	"os"
)

type RepoTree struct {
	Dir  string
	Ctxs []Repo
}

func NewTree(configDir string) (*RepoTree, error) {
	entries, err := os.ReadDir(configDir)
	switch {
	case errors.Is(err, os.ErrNotExist):
		return &RepoTree{}, nil
	case err != nil:
		return nil, err
	}

	ctxs := []Repo{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		c, err := ReadRepo(configDir, entry.Name())
		if err != nil {
			return nil, err
		}

		ctxs = append(ctxs, c)
	}

	return &RepoTree{Ctxs: ctxs}, nil
}
