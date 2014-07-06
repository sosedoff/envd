package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	VERSION = "0.1.0"
)

var options struct {
	Path  string
	Host  string
	Port  int
	Token string
	Auth  bool
}

var services []Service

func main() {
	flag.StringVar(&options.Path, "c", "", "Path to config directory")
	flag.StringVar(&options.Host, "h", "0.0.0.0", "Host to bind to")
	flag.IntVar(&options.Port, "p", 3050, "Port to listen on")
	flag.StringVar(&options.Token, "t", "", "Authentication token")

	flag.Parse()

	if options.Path == "" {
		options.Path = "./config"
	}

	// Load token from environment variable if not set
	if options.Token == "" {
		options.Token = os.Getenv("TOKEN")
	}

	// Do not require authentication if token is not set
	if options.Token == "" {
		options.Auth = false
	} else {
		options.Auth = true
	}

	var err error
	services, err = readServices(options.Path)

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("envd v%s\n", VERSION)
	fmt.Println("config path:", options.Path)
	fmt.Println("services detected:", len(services))

	startServer()
}
