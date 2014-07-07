package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Environment struct {
	Name  string
	Keys  []Key
	Hosts []string
	Token string
}

func (env Environment) ToString() string {
	lines := []string{}

	for _, key := range env.Keys {
		lines = append(lines, key.String())
	}

	return strings.Join(lines, "\n")
}

func readHosts(path string) []string {
	hosts := []string{}

	if fileExists(path) {
		str, err := ioutil.ReadFile(path)

		if err == nil {
			for _, host := range strings.Split(string(str), "\n") {
				hosts = append(hosts, strings.TrimSpace(host))
			}
		}
	}

	return hosts
}

func readToken(path string) string {
	token := ""
	str, err := ioutil.ReadFile(path)

	if err == nil {
		token = strings.TrimSpace(string(str))
	}

	return token
}

func readEnvironments(path string) []Environment {
	envs := []Environment{}
	dirs, err := getDirs(path)

	if err != nil {
		return envs
	}

	for _, dir := range dirs {
		keysDir := fmt.Sprintf("%s/%s", path, dir)
		hostsPath := fmt.Sprintf("%s/%s.hosts", path, dir)
		tokenPath := fmt.Sprintf("%s/%s.token", path, dir)

		env := Environment{
			Name:  dir,
			Keys:  readKeys(keysDir),
			Hosts: readHosts(hostsPath),
			Token: readToken(tokenPath),
		}

		envs = append(envs, env)
	}

	return envs
}
