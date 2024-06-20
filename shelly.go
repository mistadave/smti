package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func startMqttConsumer() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.1.110:1883")
	opts.SetClientID("shelly-consumer")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.SetUsername("")
	opts.SetPassword("")
	opts.SetCleanSession(true)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("shellies/+/sensor/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	if token := client.Subscribe("test/+/sensor/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		time.Sleep(1 * time.Second)
		// Do nothing
	}
}
