package main

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mistadave/smti/queue"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	queue.SENSORMQ.Channel <- queue.Data{ID: msg.Topic(), Value: msg.Payload()}
}

func startMqttConsumer(mqttBroker string) {
	opts := mqtt.NewClientOptions().AddBroker(mqttBroker)
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

}
