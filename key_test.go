package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_String(t *testing.T) {
	key := Key{Name: "foo", Value: "bar"}

	assert.Equal(t, key.String(), "FOO=bar")
}

func Test_readKeys(t *testing.T) {
	keys := readKeys("./examples/myapp/production")

	assert.Equal(t, len(keys), 4)
	assert.Equal(t, keys[0].Name, "port")
	assert.Equal(t, keys[0].Value, "5000")
	assert.Equal(t, keys[1].Name, "rails_env")
	assert.Equal(t, keys[1].Value, "production")
	assert.Equal(t, keys[2].Name, "redis_url")
	assert.Equal(t, keys[2].Value, "redis://myhost:6379/0")
	assert.Equal(t, keys[3].Name, "s3_bucket")
	assert.Equal(t, keys[3].Value, "mybucket_production")
}

func Test_readKeysEmpty(t *testing.T) {
	keys := readKeys("./tests/app/development")
	assert.Equal(t, len(keys), 0)
}
