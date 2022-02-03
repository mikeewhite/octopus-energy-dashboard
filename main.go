package main

import (
	"log"
	"os"

	influxdb "github.com/mikeewhite/octopus-energy-dashboard/influxdb"
	octopusenergy "github.com/mikeewhite/octopus-energy-dashboard/octopusenergy"

	"github.com/joho/godotenv"
)

func main() {
	// TODO - Integrate bugsnag

	// TODO - Create a config package to encapsulate all of this config and derive it from env vars (and error if one is not set)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}

	apiKey := os.Getenv("OCTOPUS_ENERGY_API_KEY")
	if apiKey == "" {
		log.Fatal("missing OCTOPUS_ENERGY_API_KEY env var")
	}

	mpan := os.Getenv("ELECTRICITY_METER_MPAN")
	if mpan == "" {
		log.Fatal("missing ELECTRICITY_METER_MPAN env var")
	}

	serialNumber := os.Getenv("ELECTRICITY_METER_SERIAL_NUMBER")
	if serialNumber == "" {
		log.Fatal("missing ELECTRICITY_METER_SERIAL_NUMBER env var")
	}

	influxDBToken := os.Getenv("INFLUXDB_TOKEN")
	if influxDBToken == "" {
		log.Fatal("missing INFLUXDB_TOKEN env var")
	}
	influxDBClient := influxdb.New(influxDBToken)

	octopusEnergyClient := octopusenergy.New(apiKey, mpan, serialNumber)
	consumption, err := octopusEnergyClient.GetElectricityMeterConsumption()
	if err != nil {
		log.Fatalf("error on retrieving electricity meter consumption: %s", err)
	}

	// TODO - consider injecting influx client into octopus energy package so that it can record the consumption itself (using interfaces)
	influxDBClient.RecordConsumption(consumption)

	// TODO - Push data to InfluxDB
	// TODO - Create grafana board to visualize data

	// TODO - Use docker volumes so that data is persisted
	// TODO - Configure grafana to use pre-set datasource and layout on startup
}
