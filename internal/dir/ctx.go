package dir

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Ctx struct {
	Root    string
	Repo    string
	Current string
	Envs    []string
}

func CreateCtx(configDir, repo string) (Ctx, error) {
	if err := os.MkdirAll(filepath.Join(configDir, repo), 0700); err != nil {
		return Ctx{}, err
	}

	return ReadCtx(configDir, repo)
}

func ReadCtx(root, repo string) (Ctx, error) {
	c := Ctx{Root: root, Repo: repo}

	repoRoot := filepath.Join(root, repo)

	if _, err := os.Stat(repoRoot); err != nil {
		return c, err
	}

	repoCtxs, err := os.ReadDir(repoRoot)
	if err != nil {
		return c, err
	}

	for _, v := range repoCtxs {
		if v.IsDir() {
			continue
		}

		n := v.Name()
		switch {
		case strings.HasSuffix(n, ".yaml"):
			c.Envs = append(c.Envs, n[:len(n)-5])
		case n == "current":
			buf, err := os.ReadFile(filepath.Join(repoRoot, n))
			if err != nil {
				return c, err
			}
			c.Current = strings.TrimSpace(string(buf))
		}
	}

	return c, nil
}

func (c *Ctx) Path() string {
	return filepath.Join(c.Root, c.Repo)
}

func (c *Ctx) AddEnv(root, name string) error {
	if err := os.WriteFile(filepath.Join(root, c.Repo, name+".yaml"), []byte{}, 0600); err != nil {
		return err
	}

	c.Envs = append(c.Envs, name)
	if len(c.Envs) > 1 {
		return nil
	}

	return os.WriteFile(filepath.Join(root, c.Repo, "current"), []byte(strings.TrimSpace(name)), 0600)
}

// Read the current context
func (c *Ctx) Read() ([]byte, error) {
	buf, err := os.ReadFile(filepath.Join(c.Path(), "current"))
	if err != nil {
		return nil, fmt.Errorf("failed reading current ctx: %w", err)
	}

	s := string(buf)
	for _, v := range c.Envs {
		if v == s {
			return os.ReadFile(filepath.Join(c.Path(), v+".yaml"))
		}
	}

	return nil, fmt.Errorf("env %s not found but was what was in current", s)
}
