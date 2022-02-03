package influxdb

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/domain"

	octopusenergy "github.com/mikeewhite/octopus-energy-dashboard/octopusenergy"
)

// TODO - These should be set via config
const bucketName = "consumption"
const orgName = "energy"

// Client for interfacing with InfluxDB
type Client struct {
	influxClient influxdb2.Client
}

// New creates a new client
func New(token string) *Client {
	url := "http://localhost:8086" // TODO - this should be sourced from a config package that defines it properly
	influxClient := influxdb2.NewClient(url, token)

	// check InfluxDB is running
	health, err := influxClient.Health(context.Background())
	if err != nil {
		panic("error on performing InfluxDB health check")
	}
	if health.Status != domain.HealthCheckStatusPass {
		panic("InfluxDB health check failed")
	}

	return &Client{
		influxClient: influxClient,
	}
}

// RecordConsumption TODO
func (c *Client) RecordConsumption(consumption *octopusenergy.Consumption) {
	writeAPI := c.influxClient.WriteAPI(orgName, bucketName)
	for _, result := range consumption.Results {
		p := influxdb2.NewPointWithMeasurement("consumption").
			AddTag("unit", "kwh").
			AddField("consumption", result.Consumption).
			SetTime(result.IntervalStart.UTC()) // TODO - Derive time from consumption
		writeAPI.WritePoint(p)
	}
	writeAPI.Flush()
}
