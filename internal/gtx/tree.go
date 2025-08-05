package gtx

import (
	"errors"
	"os"
)

type Tree struct {
	Dir  string
	Ctxs []Ctx
}

func NewTree(configDir string) (*Tree, error) {
	entries, err := os.ReadDir(configDir)
	switch {
	case errors.Is(err, os.ErrNotExist):
		return &Tree{}, nil
	case err != nil:
		return nil, err
	}

	ctxs := []Ctx{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		c, err := ReadCtx(configDir, entry.Name())
		if err != nil {
			return nil, err
		}

		ctxs = append(ctxs, c)
	}

	return &Tree{Ctxs: ctxs}, nil
}
