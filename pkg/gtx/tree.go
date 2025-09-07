package gtx

import (
	"errors"
	"os"
)

type RepoTree struct {
	// Directory root location
	Dir string
	// All repos
	Repos []Repo
}

// Gets all configuration information in the filesystem
// located at configDir
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

	return &RepoTree{Repos: ctxs}, nil
}
