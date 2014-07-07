package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isJunkFile(t *testing.T) {
	assert.Equal(t, isJunkFile(".DS_Store"), true)
	assert.Equal(t, isJunkFile(".gitkeep"), true)
	assert.Equal(t, isJunkFile(".gitignore"), true)
	assert.Equal(t, isJunkFile("file"), false)
}

func Test_fileExists(t *testing.T) {
	assert.Equal(t, fileExists("/tmp"), true)
	assert.Equal(t, fileExists("/tmp/foobar"), false)
}

func Test_getDirs(t *testing.T) {
	dirs, err := getDirs("./examples/myapp")

	assert.Equal(t, err, nil)
	assert.Equal(t, len(dirs), 2)
	assert.Equal(t, dirs[0], "production")
	assert.Equal(t, dirs[1], "staging")
}

func Test_getDirsFailure(t *testing.T) {
	_, err := getDirs("./examples/myapp2")

	assert.NotEqual(t, err, nil)
}

func Test_getDirFiles(t *testing.T) {
	files, err := getDirFiles("./examples/myapp/production")

	assert.Equal(t, err, nil)
	assert.Equal(t, len(files), 4)
	assert.Equal(t, files[0], "port")
	assert.Equal(t, files[1], "rails_env")
	assert.Equal(t, files[2], "redis_url")
	assert.Equal(t, files[3], "s3_bucket")
}

func Test_getDirFilesEmpty(t *testing.T) {
	files, err := getDirFiles("./tests/empty")

	assert.Equal(t, err, nil)
	assert.Equal(t, len(files), 0)
}

func Test_getDirFilesFailure(t *testing.T) {
	_, err := getDirFiles("./tests/emptyfoo")

	assert.NotEqual(t, err, nil)
}
