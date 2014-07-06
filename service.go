package main

import (
	"fmt"
)

type Service struct {
	Name         string
	Environments []Environment
}

func (service Service) EnvironmentNames() []string {
	names := []string{}

	for _, env := range service.Environments {
		names = append(names, env.Name)
	}

	return names
}

func readServices(path string) (result []Service, err error) {
	dirs, err := getDirs(path)

	if err != nil {
		return
	}

	for _, name := range dirs {
		dir := fmt.Sprintf("%s/%s", path, name)
		service := Service{Name: name, Environments: readEnvironments(dir)}

		result = append(result, service)
	}

	err = nil
	return
}

func getService(name string) (svc Service, err error) {
	err = fmt.Errorf("Invalid service")

	for _, service := range services {
		if service.Name == name {
			svc = service
			err = nil
			return
		}
	}

	return
}

func getEnvironment(serviceName string, envName string) (env Environment, err error) {
	service, err := getService(serviceName)

	if err != nil {
		return
	}

	err = fmt.Errorf("Invalid environment")

	for _, e := range service.Environments {
		if e.Name == envName {
			env = e
			err = nil
			return
		}
	}

	return
}
