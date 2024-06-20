package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting Shelly mqtt Consumer")
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	fmt.Println("Environment:", *environment)
	startMqttConsumer()
}
