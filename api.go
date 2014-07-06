package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// Returns client authentication token from header or url params
func getClientToken(c *gin.Context) string {
	// Try fetching auth token from headers first
	token := c.Req.Header.Get("Token")

	// Try to fetch from url param if blank
	if token == "" {
		token = c.Req.URL.Query()["token"][0]
	}

	return token
}

func RequireAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getClientToken(c)

		if token != options.Token {
			c.Abort(401)
		}
	}
}

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

	if options.Auth {
		fmt.Println("authentication enabled")
		api.Use(RequireAuthToken())
	} else {
		fmt.Println("authentication disabled")
	}

	api.GET("/", renderAvailableServices)
	api.GET("/:service", renderServiceEnvironments)
	api.GET("/:service/:env", renderServiceEnvironment)
	api.POST("/reload", renderReloadServices)

	fmt.Println("starting server on port", options.Port)
	api.Run(fmt.Sprintf(":%d", options.Port))
}
