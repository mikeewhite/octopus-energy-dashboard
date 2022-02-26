package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config TODO
type Config struct {
	OctopusEnergy OctopusEnergy
	InfluxDB      InfluxDB
}

// OctopusEnergy TODO
type OctopusEnergy struct {
	APIKey                   string
	ElectricityMeterMPAN     string
	ElectricityMeterSerialNo string
}

// InfluxDB TODO
type InfluxDB struct {
	BucketName string
	OrgName    string
	Token      string
	URL        string
}

// New TODO
func New() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("error loading .env file: %s", err))
	}

	return &Config{
		OctopusEnergy: OctopusEnergy{
			APIKey:                   resolveEnvVar("OCTOPUS_ENERGY_API_KEY"),
			ElectricityMeterMPAN:     resolveEnvVar("ELECTRICITY_METER_MPAN"),
			ElectricityMeterSerialNo: resolveEnvVar("ELECTRICITY_METER_SERIAL_NUMBER"),
		},
		InfluxDB: InfluxDB{
			BucketName: "consumption",
			OrgName:    "energy",
			Token:      resolveEnvVar("INFLUXDB_TOKEN"),
			URL:        "http://localhost:8086",
		},
	}
}

func resolveEnvVar(envVarName string) string {
	value := os.Getenv(envVarName)
	if value == "" {
		panic(fmt.Sprintf("missing %s env var", envVarName))
	}
	return value
}
