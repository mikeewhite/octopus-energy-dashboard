# 🐙 Octopus Energy Dashboard

## Usage
1. Create a `.env` file in the root directory with the following content:
```bash
OCTOPUS_ENERGY_API_KEY=<YOUR-API-KEY>
ELECTRICITY_METER_MPAN=<YOUR-MPAN-NUMBER>
ELECTRICITY_METER_SERIAL_NUMBER=<YOUR-SERIAL-NUMBER>
INFLUXDB_USERNAME=admin
INFLUXDB_PASSWORD=admin1234
INFLUXDB_TOKEN=<YOUR-INFLUXDB-TOKEN>
```
> You can get these details from https://octopus.energy/dashboard/developer/

2. Start the dashboard with
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