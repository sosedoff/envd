package main

import (
	"fmt"
	"strings"
)

type Environment struct {
	Name string
	Keys []Key
}

func (env Environment) ToString() string {
	lines := []string{}

	for _, key := range env.Keys {
		lines = append(lines, key.String())
	}

	return strings.Join(lines, "\n")
}

func readEnvironments(path string) []Environment {
	envs := []Environment{}
	dirs, err := getDirs(path)

	if err != nil {
		return envs
	}

	for _, dir := range dirs {
		keysDir := fmt.Sprintf("%s/%s", path, dir)
		env := Environment{Name: dir, Keys: readKeys(keysDir)}
		envs = append(envs, env)
	}

	return envs
}
