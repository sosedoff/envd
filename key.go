package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Key struct {
	Name  string
	Value string
}

func (key Key) String() string {
	return fmt.Sprintf("%s=%s", strings.ToUpper(key.Name), key.Value)
}

func readKeys(path string) []Key {
	keys := []Key{}
	files, _ := getDirFiles(path)

	for _, file := range files {
		content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, file))

		if err != nil {
			continue
		}

		keys = append(keys, Key{
			Name:  file,
			Value: strings.TrimSpace(string(content)),
		})
	}

	return keys
}
