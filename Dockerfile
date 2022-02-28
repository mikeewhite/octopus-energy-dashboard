# Build source code into a runnable binary
FROM golang:1.17.7-alpine3.15 as builder

# Disbale CGO otherwise the built binary will not run in Ubuntu due to linked files
ENV CGO_ENABLED=0

# Environment varibales are not passed to later build phases so we have to set them as global build args
# See https://github.com/moby/moby/issues/37345
ARG OCTOPUS_ENERGY_API_KEY
ARG ELECTRICITY_METER_MPAN
ARG ELECTRICITY_METER_SERIAL_NUMBER
ARG INFLUXDB_USERNAME
ARG INFLUXDB_PASSWORD
ARG INFLUXDB_TOKEN

COPY . /app
WORKDIR /app
RUN go build -o data-collector

# Run binary as a cron job
FROM ubuntu:latest as cron

ARG OCTOPUS_ENERGY_API_KEY
ENV OCTOPUS_ENERGY_API_KEY=$OCTOPUS_ENERGY_API_KEY
ARG ELECTRICITY_METER_MPAN
ENV ELECTRICITY_METER_MPAN=$ELECTRICITY_METER_MPAN
ARG ELECTRICITY_METER_SERIAL_NUMBER
ENV ELECTRICITY_METER_SERIAL_NUMBER=$ELECTRICITY_METER_SERIAL_NUMBER
ARG INFLUXDB_USERNAME
ENV INFLUXDB_USERNAME=$INFLUXDB_USERNAME
ARG INFLUXDB_PASSWORD
ENV INFLUXDB_PASSWORD=$INFLUXDB_PASSWORD
ARG INFLUXDB_TOKEN
ENV INFLUXDB_TOKEN=$INFLUXDB_TOKEN

RUN apt-get update && apt-get -y install cron

COPY --from=builder /app/data-collector /data-collector
RUN chmod 0744 /data-collector

ADD crontab /etc/cron.d/simple-cron
RUN chmod 0644 /etc/cron.d/simple-cron
RUN crontab /etc/cron.d/simple-cron

RUN touch /var/log/cron.log
CMD cron && tail -f /var/log/cron.log