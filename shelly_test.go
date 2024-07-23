package main

import (
	"testing"
)

func TestParseTopic(t *testing.T) {
	t.Run("Valid topic", func(t *testing.T) {
		topic := "shellies/shellyht-F2B61C/sensor/temperature"
		expectedDevice := "shellyht-F2B61C"
		expectedSensor := "temperature"

		device, sensor := parseTopic(topic)

		if device != expectedDevice {
			t.Errorf("Expected device to be %s, but got %s", expectedDevice, device)
		}

		if sensor != expectedSensor {
			t.Errorf("Expected sensor to be %s, but got %s", expectedSensor, sensor)
		}
	})

	t.Run("Valid topic with multiple sensors", func(t *testing.T) {
		topic := "shellies/shellybt-F2B61C/sensor/humidity"
		expectedDevice := "shellybt-F2B61C"
		expectedSensor := "humidity"

		device, sensor := parseTopic(topic)

		if device != expectedDevice {
			t.Errorf("Expected device to be %s, but got %s", expectedDevice, device)
		}

		if sensor != expectedSensor {
			t.Errorf("Expected sensor to be %s, but got %s", expectedSensor, sensor)
		}
	})

	t.Run("Invalid topic", func(t *testing.T) {
		topic := "invalid_topic"

		device, sensor := parseTopic(topic)

		if device != "" {
			t.Errorf("Expected device to be empty, but got %s", device)
		}

		if sensor != "" {
			t.Errorf("Expected sensor to be empty, but got %s", sensor)
		}
	})
}
