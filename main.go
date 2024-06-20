package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/reMarkable/envconfig/v2"
)

type Config struct {
	Environment string `envconfig:"ENVIRONMENT" default:"development"`
	MqttBroker  string `envconfig:"MQTT_BROKER" default:"tcp://localhost:1883" required:"true"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found to load")
		// os.Exit(1)
	}
}

func main() {
	fmt.Println("Starting Shelly mqtt Consumer")

	var config Config
	if err := envconfig.Process("myapp", &config); err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}
	fmt.Println("configs:", config)
	fmt.Println("environment:", config.Environment)
	startMqttConsumer(config.MqttBroker)
}
