package main

import (
	"fmt"
	"os"
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

	fmt.Println("Services detected:", len(services))
	fmt.Println("Starting server on port", options.Port)

	startServer()
}
