package main

import (
	"log"

	"github.com/mikeewhite/octopus-energy-dashboard/pkg/config"
	influxdb "github.com/mikeewhite/octopus-energy-dashboard/pkg/influxdb"
	octopusenergy "github.com/mikeewhite/octopus-energy-dashboard/pkg/octopusenergy"
)

func main() {
	cfg := config.New()
	influxDBClient := influxdb.New(cfg.InfluxDB)
	octopusEnergyClient := octopusenergy.New(cfg.OctopusEnergy)
	consumption, err := octopusEnergyClient.GetElectricityMeterConsumption()
	if err != nil {
		log.Fatalf("error on retrieving electricity meter consumption: %s", err)
	}
	influxDBClient.RecordConsumption(consumption)
}
