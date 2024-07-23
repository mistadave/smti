package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mistadave/smti/queue"
)

func parseTopic(topic string) (string, string) {
	tokens := strings.Split(topic, "/")
	if len(tokens) >= 4 {
		return tokens[1], tokens[3]
	}
	return "", ""
}

type KeyValue struct {
	Key   string
	Value float64
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	id, key := parseTopic(msg.Topic())
	value, err := strconv.ParseFloat(string(msg.Payload()), 32)
	if err != nil {
		fmt.Println("Error parsing payload:", err)
		return
	}
	queue.SENSORMQ.Channel <- queue.Data{ID: id, Value: KeyValue{key, value}}
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
