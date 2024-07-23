package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mistadave/smti/config"
	"github.com/mistadave/smti/db"
	"github.com/reMarkable/envconfig/v2"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found to load")
		// os.Exit(1)
	}
}

func main() {
	fmt.Println("Starting Shelly mqtt Consumer")
	var config config.Config
	if err := envconfig.Process("", &config); err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	fmt.Println("configs:", config)
	fmt.Println("environment:", config.Environment)
	db.Init(config)
	startMqttConsumer(config.MqttBroker)
	db.StartInfluxConsumer()
	for {
		select {}
	}
}
