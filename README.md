# 🐙 Octopus Energy Dashboard

Prototype dashboard for viewing electricity usage consumed from the [Octopus Energy API](https://developer.octopus.energy/docs/api/#consumption) using Golang (for data collection), InfluxDB (for data storage), and Grafana (for visualization).

## Usage
1. Create a `.env` file in the root directory with the following content:
```bash
# Octopus energy config (get your values from https://octopus.energy/dashboard/developer/)
OCTOPUS_ENERGY_API_KEY=<YOUR-API-KEY>
ELECTRICITY_METER_MPAN=<YOUR-MPAN-NUMBER>
ELECTRICITY_METER_SERIAL_NUMBER=<YOUR-SERIAL-NUMBER>

# InfluxDB config
INFLUXDB_USERNAME=admin
INFLUXDB_PASSWORD=admin1234
INFLUXDB_TOKEN=O-mfkVnkRYlueeyffw8q0T_K2Cf4TJMtFGlZaZoFxG-v80ZhvWSGZyJwMaRrIAIHWtA6pZ_bDQCwTvApccFcVw==
```

2. Start the dashboard with:
```bash
docker-compose up -d
```

3. Run `main.go` to collect data.

4. Create [DBRP mappings](https://docs.influxdata.com/influxdb/v2.0/query-data/influxql/#map-unmapped-buckets):
```bash
docker-compose exec influxdb bash
influx bucket list # to determine ID of 'consumption bucket' - see https://docs.influxdata.com/influxdb/v2.1/organizations/buckets/view-buckets
influx v1 dbrp create --db consumption --rp infinite --bucket-id <bucket-id> --default
```

5. Visit http://localhost:3001 and navigate to 'Consumption (InfluxQL)' dashboard.