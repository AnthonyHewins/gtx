package gtx

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
)

func SetCtx(repoDir, env string) error {
	entries, err := os.ReadDir(repoDir)
	if err != nil {
		return err
	}

	for _, v := range entries {
		if v.Name() == env+".yaml" {
			return os.WriteFile(filepath.Join(repoDir, "current"), []byte(env), 0600)
		}
	}

	return fmt.Errorf("env %s not found in repo %s", env, repoDir)
}

func Read[X any](repoDir string) (X, error) {
	current, err := os.ReadFile(filepath.Join(repoDir, "current"))
	if err != nil {
		var x X
		return x, err
	}

	return ReadEnv[X](repoDir, strings.TrimSpace(string(current)))
}

func ReadEnv[X any](repoDir, env string) (X, error) {
	var x X

	path := filepath.Join(repoDir, env+".yaml")
	buf, err := os.ReadFile(path)
	if err != nil {
		return x, err
	}

	return x, yaml.Unmarshal(buf, &x)
}
