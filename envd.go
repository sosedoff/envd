package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const (
	VERSION = "0.3.0"
)

var options struct {
	Path  string
	Host  string
	Port  int
	Token string
	Auth  bool
}

var services []Service

func initOptions() {
	var printVersion bool

	flag.StringVar(&options.Path, "c", "", "Path to config directory")
	flag.StringVar(&options.Host, "h", "0.0.0.0", "Host to bind to")
	flag.IntVar(&options.Port, "p", 3050, "Port to listen on")
	flag.StringVar(&options.Token, "t", "", "Authentication token")
	flag.BoolVar(&printVersion, "v", false, "Print version")

	flag.Parse()

	if printVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if options.Path == "" {
		fmt.Println("Please specify -c option")
		os.Exit(1)
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
}

func reloadServices() {
	newServices, err := readServices(options.Path)

	if err != nil {
		return
	}

	services = newServices
}

func setupReload() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func() {
		for sig := range c {
			fmt.Println("Reloading configuration...", sig)
			reloadServices()
		}
	}()
}

func main() {
	initOptions()
	setupReload()

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
