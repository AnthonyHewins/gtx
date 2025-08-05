package gtx

import (
	"os"
	"path/filepath"
	"strings"
)

type Ctx struct {
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

func (c *Ctx) AddEnv(root, name string) error {
	if err := os.WriteFile(filepath.Join(root, c.Repo, name+".yaml"), []byte{}, 0600); err != nil {
		return err
	}

	c.Envs = append(c.Envs, name)
	if len(c.Envs) > 1 {
		return nil
	}

	return os.WriteFile(filepath.Join(root, c.Repo, "current"), []byte(name), 0600)
}

func ReadCtx(root, repo string) (Ctx, error) {
	c := Ctx{Repo: repo}

	repoRoot := filepath.Join(root, repo)
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
