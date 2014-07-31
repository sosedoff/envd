package main

import (
	"io/ioutil"
	"os"
)

var skipFiles = [...]string{
	".DS_Store",
	".gitkeep",
	".gitignore",
	".git",
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isJunkFile(name string) bool {
	for _, file := range skipFiles {
		if file == name {
			return true
		}
	}

	return false
}

func getDirs(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	dirs := []string{}

	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}

	return dirs, nil
}

func getDirFiles(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	result := []string{}

	for _, file := range files {
		if !file.IsDir() && !isJunkFile(file.Name()) {
			result = append(result, file.Name())
		}
	}

	return result, nil
}
