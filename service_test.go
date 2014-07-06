package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EnvironmentNames(t *testing.T) {
	service := Service{
		Name: "foo",
		Environments: []Environment{
			Environment{Name: "production"},
			Environment{Name: "staging"},
		},
	}

	assert.Equal(t, service.EnvironmentNames(), []string{"production", "staging"})
}

func Test_EnvironmentNamesEmpty(t *testing.T) {
	service := Service{Name: "foo"}

	assert.Equal(t, service.EnvironmentNames(), []string{})
}

func Test_readServices(t *testing.T) {
	result, err := readServices("./examples")

	assert.NoError(t, err)
	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0].Name, "myapp")
	assert.Equal(t, len(result[0].Environments), 2)
}

func Test_readServicesEmpty(t *testing.T) {
	result, err := readServices("./examples/myapp/production")

	assert.NoError(t, err)
	assert.Equal(t, len(result), 0)
}

func Test_readServicesFail(t *testing.T) {
	_, err := readServices("./examples2")

	assert.Error(t, err)
}

func Test_getService(t *testing.T) {
	services = []Service{Service{Name: "myapp"}}
	result, err := getService("myapp")

	assert.NoError(t, err)
	assert.IsType(t, Service{}, result)
	assert.Equal(t, result.Name, "myapp")
}

func Test_getServiceFail(t *testing.T) {
	services = []Service{Service{Name: "myapp"}}
	_, err := getService("myapp2")

	assert.Error(t, err)
}

func Test_getEnvironment(t *testing.T) {
	services = []Service{
		Service{
			Name: "myapp",
			Environments: []Environment{
				Environment{Name: "production"},
			},
		},
	}

	result, err := getEnvironment("myapp", "production")

	assert.NoError(t, err)
	assert.IsType(t, Environment{}, result)
	assert.Equal(t, result.Name, "production")
}

func Test_getEnvironmentFail(t *testing.T) {
	services = []Service{
		Service{
			Name: "myapp",
			Environments: []Environment{
				Environment{Name: "production"},
			},
		},
	}

	_, err := getEnvironment("myapp", "staging")

	assert.Error(t, err)
}
