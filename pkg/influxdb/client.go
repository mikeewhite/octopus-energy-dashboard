package influxdb

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/domain"

	"github.com/mikeewhite/octopus-energy-dashboard/pkg/config"
	octopusenergy "github.com/mikeewhite/octopus-energy-dashboard/pkg/octopusenergy"
)

// Client for interfacing with InfluxDB
type Client struct {
	bucketName   string
	orgName      string
	influxClient influxdb2.Client
}

// New creates a new client
func New(cfg config.InfluxDB) *Client {
	influxClient := influxdb2.NewClient(cfg.URL, cfg.Token)

	// check InfluxDB is running
	health, err := influxClient.Health(context.Background())
	if err != nil {
		panic("error on performing InfluxDB health check")
	}
	if health.Status != domain.HealthCheckStatusPass {
		panic("InfluxDB health check failed")
	}

	return &Client{
		bucketName:   cfg.BucketName,
		orgName:      cfg.OrgName,
		influxClient: influxClient,
	}
}

// RecordConsumption TODO
func (c *Client) RecordConsumption(consumption *octopusenergy.Consumption) {
	writeAPI := c.influxClient.WriteAPI(c.orgName, c.bucketName)
	for _, result := range consumption.Results {
		p := influxdb2.NewPointWithMeasurement("consumption").
			AddTag("unit", "kwh").
			AddField("consumption", result.Consumption).
			SetTime(result.IntervalStart.UTC())
		writeAPI.WritePoint(p)
	}
	writeAPI.Flush()
}
