package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"strings"
)

// Returns client authentication token from header or url params
func getClientToken(c *gin.Context) string {
	// Try fetching auth token from headers first
	token := c.Req.Header.Get("Token")

	// Try to fetch from url param if blank
	if token == "" {
		if len(c.Req.URL.Query()["token"]) > 0 {
			token = c.Req.URL.Query()["token"][0]
		}
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

	// Respond with 400 if service does not exist
	// Todo: maybe respond with 404
	if err != nil {
		c.String(400, err.Error()+"\n")
		return
	}

	// Get remote IP address
	host, _, err := net.SplitHostPort(c.Req.RemoteAddr)

	if err != nil {
		c.String(400, err.Error()+"\n")
		return
	}

	// Check if environment has allowed hosts. Localhost is allowed.
	if host != "::1" && len(environment.Hosts) > 0 {
		// Reject requests from non-whitelisted hosts
		if environment.HostEnabled(host) == false {
			c.String(401, "Restricted\n")
			return
		}
	}

	// Check if environment has access token
	// Environment tokens DO NOT work if global access token is set
	// Global token is set with -t flag or via TOKEN env variable
	if !options.Auth && environment.Token != "" {
		token := getClientToken(c)

		if token != environment.Token {
			c.Abort(401)
			return
		}
	}

	c.String(200, environment.ToString()+"\n")
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

	host := fmt.Sprintf("%s:%d", options.Host, options.Port)

	fmt.Println("starting server on", host)
	api.Run(host)
}
