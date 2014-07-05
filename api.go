package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func renderAvailableServices(c *gin.Context) {
	names := []string{}

	for _, svc := range services {
		names = append(names, svc.Name)
	}

	c.String(200, strings.Join(names, "\n")+"\n")
}

func renderServiceEnvironments(c *gin.Context) {
	serviceName := c.Params.ByName("service")

	service, err := getService(serviceName)
	if err != nil {
		c.String(400, err.Error()+"\n")
		return
	}

	names := strings.Join(service.EnvironmentNames(), "\n") + "\n"
	c.String(200, names)
}

func renderServiceEnvironment(c *gin.Context) {
	serviceName := c.Params.ByName("service")
	envName := c.Params.ByName("env")

	environment, err := getEnvironment(serviceName, envName)

	if err != nil {
		c.String(400, err.Error()+"\n")
		return
	}

	c.String(200, environment.ToString()+"\n")
}

func renderReloadServices(c *gin.Context) {
	new_services, err := readServices()

	if err != nil {
		c.String(400, err.Error()+"\n")
		return
	}

	// Replace current configuration
	services = new_services

	c.String(200, "OK\n")
}

func startServer() {
	api := gin.Default()

	api.GET("/", renderAvailableServices)
	api.GET("/:service", renderServiceEnvironments)
	api.GET("/:service/:env", renderServiceEnvironment)
	api.POST("/reload", renderReloadServices)

	api.Run(fmt.Sprintf(":%d", options.Port))
}
