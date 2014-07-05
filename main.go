package main

import (
	"fmt"
	"os"
)

const (
	VERSION = "0.1.0"
)

var options struct {
	Path string
	Port int
}

var services []Service

func main() {
	options.Path = "./config"
	options.Port = 5000

	var err error
	services, err = readServices()

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("envd v%s\n\n", VERSION)
	fmt.Println("Services detected:", len(services))
	fmt.Println("Starting server on port", options.Port)

	startServer()
}
