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
	token := c.Request.Header.Get("Token")

	// Try to fetch from url param if blank
	if token == "" {
		if len(c.Request.URL.Query()["token"]) > 0 {
			token = c.Request.URL.Query()["token"][0]
		}
	}

	return token
}

// Returns a list of all available services
func renderAvailableServices(c *gin.Context) {
	// Check if authentication is enabled
	if options.Auth {
		token := getClientToken(c)

		if token != options.Token {
			c.String(401, "Invalid token")
			c.Abort()
			return
		}
	}

	names := []string{}
	for _, svc := range services {
		names = append(names, svc.Name)
	}

	c.String(200, strings.Join(names, "\n")+"\n")
}

// Returns a list of all service environments
func renderServiceEnvironments(c *gin.Context) {
	// Check if authentication is enabled
	if options.Auth {
		token := getClientToken(c)

		if token != options.Token {
			c.String(401, "Invalid token")
			c.Abort()
			return
		}
	}

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
	host, _, err := net.SplitHostPort(c.Request.RemoteAddr)

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

	// Fetch token from url param or from the header
	token := getClientToken(c)

	// Validate environment token if its set, otherwise check agains global token
	if environment.Token != "" {
		if token != environment.Token {
			c.String(401, "Restricted\n")
			return
		}
	} else {
		if options.Auth && token != options.Token {
			c.String(401, "Restricted\n")
			return
		}
	}

	c.String(200, environment.ToString()+"\n")
}

func startServer() {
	host := fmt.Sprintf("%s:%d", options.Host, options.Port)
	api := gin.Default()

	api.GET("/", renderAvailableServices)
	api.GET("/:service", renderServiceEnvironments)
	api.GET("/:service/:env", renderServiceEnvironment)

	if options.Auth {
		fmt.Println("authentication enabled")
	}

	fmt.Println("starting server on", host)
	api.Run(host)
}
