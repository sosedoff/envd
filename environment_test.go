package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ToStringWithoutKeys(t *testing.T) {
	env := Environment{Name: "production"}
	assert.Equal(t, env.ToString(), "")
}

func Test_ToStringWithKeys(t *testing.T) {
	env := Environment{
		Name: "production",
		Keys: []Key{
			Key{Name: "foo", Value: "bar"},
			Key{Name: "hello", Value: "world"},
		},
	}

	assert.Equal(t, env.ToString(), "FOO=bar\nHELLO=world")
}

func Test_readEnvironments(t *testing.T) {
	envs := readEnvironments("./examples/myapp")

	assert.Equal(t, len(envs), 2)

	assert.Equal(t, envs[0].Name, "production")
	assert.Equal(t, len(envs[0].Keys), 4)
	assert.Equal(t, len(envs[0].Hosts), 3)
	assert.Equal(t, envs[0].Token, "sampletoken")

	assert.Equal(t, envs[1].Name, "staging")
	assert.Equal(t, len(envs[1].Keys), 4)
	assert.Equal(t, len(envs[1].Hosts), 0)
	assert.Equal(t, envs[1].Token, "")
}
