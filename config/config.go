package config

type Config struct {
	Environment  string `envconfig:"ENVIRONMENT" default:"development"`
	MqttBroker   string `envconfig:"MQTT_BROKER" default:"tcp://localhost:1883" required:"true"`
	InfluxUrl    string `envconfig:"INFLUX_URL" default:"http://localhost:8086" required:"true"`
	InfluxToken  string `envconfig:"INFLUX_TOKEN" required:"true"`
	InfluxBucket string `envconfig:"INFLUX_BUCKET" required:"true"`
	InfluxOrg    string `envconfig:"INFLUX_ORG" default:"schloss" required:"true"`
}
