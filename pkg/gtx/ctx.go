package gtx

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Repo is a data structure with all the contexts available for a repository
// and where it's located at the home system. It represents a folder location on
// your file system, and various files on your filesystem
type Repo struct {
	// Root directory of all repos, likely $HOME/.config/gtx unless configured otherwise
	Root string
	// Repo name, which is joined together with Root
	Repo string
	// Current active repo that will be loaded
	Current string
	// List of all envs for the dir
	Envs []string
}

// Create a new repo
func CreateRepo(configDir, repo string) (Repo, error) {
	if err := os.MkdirAll(filepath.Join(configDir, repo), 0700); err != nil {
		return Repo{}, err
	}

	return ReadRepo(configDir, repo)
}

// Reads the filesystem for a repo
func ReadRepo(root, repo string) (Repo, error) {
	c := Repo{Root: root, Repo: repo}

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

// Gets the repo path doing a filepath.Join of Root and Repo,
// giving you the path of the repo's contexts and the current context
func (c *Repo) Path() string {
	return filepath.Join(c.Root, c.Repo)
}

// Add a new context to the repo
func (c *Repo) AddCtx(root, name string) error {
	if err := os.WriteFile(filepath.Join(root, c.Repo, name+".yaml"), []byte{}, 0600); err != nil {
		return err
	}

	if c.Envs = append(c.Envs, name); len(c.Envs) > 1 {
		return nil
	}

	return os.WriteFile(filepath.Join(root, c.Repo, "current"), []byte(strings.TrimSpace(name)), 0600)
}

// Read the current context
func (c *Repo) Read() ([]byte, error) {
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
