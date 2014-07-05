package main

import (
	"io/ioutil"
)

func isJunkFile(name string) bool {
	if name == ".DS_Store" {
		return true
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
