package db

import (
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/mistadave/smti/config"
	"github.com/mistadave/smti/queue"
)

var influxClient influxdb2.Client
var writeAPI api.WriteAPI

func Init(config config.Config) {
	influxClient = influxdb2.NewClient(config.InfluxUrl, config.InfluxToken)
	writeAPI = influxClient.WriteAPI(config.InfluxOrg, config.InfluxBucket)
}

func StartInfluxConsumer() {
	for {
		select {
		case data := <-queue.SENSORMQ.Channel:
			fmt.Println("Received data from queue:", data)
			p := influxdb2.NewPointWithMeasurement("sensor").
				AddTag("id", data.ID).
				AddField("value", data.Value).
				SetTime(time.Now())
			for k, v := range data.Tags {
				p.AddTag(k, v)
			}
			writeAPI.WritePoint(p)
		}
	}
}

func CloseInfluxClient() {
	influxClient.Close()
}
