package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func generateRandomFloat(start int, end int) float64 {
	min := float64(start)
	max := float64(end)
	return min + rand.Float64()*(max-min)
}

func sendMsgByIdKey(client mqtt.Client, id string, key string) {
	value := 0.0
	if key == "temperature" {
		value = generateRandomFloat(20, 30)
	} else if key == "humidity" {
		value = generateRandomFloat(40, 60)
	} else {
		value = generateRandomFloat(1, 10)
	}
	client.Publish(fmt.Sprintf("shellies/%s/sensor/%s", id, key), 0, false, fmt.Sprintf("%f", value))
}

func main() {
	fmt.Println("Starting mqtt Producer")
	mqttBroker := flag.String("mqttBroker", "tcp://localhost:1883", "MQTT broker address")
	flag.Parse()

	opts := mqtt.NewClientOptions().AddBroker(*mqttBroker)
	opts.SetClientID("shelly-producer")
	opts.SetUsername("")
	opts.SetPassword("")
	opts.SetCleanSession(true)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	minTicker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			sendMsgByIdKey(client, "shelly1", "temperature")
			sendMsgByIdKey(client, "shelly1", "humidity")
			sendMsgByIdKey(client, "shelly2", "temperature")
			sendMsgByIdKey(client, "shelly2", "humidity")
		case <-minTicker.C:
			fmt.Println("After 1 min")
		}
	}

}
