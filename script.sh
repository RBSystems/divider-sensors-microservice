#!/usr/bin/env bash
echo Enter the DSP microservice address for these sensors: read DSPaddress
"DSP_MICROSERVICE_ADDRESS=$DSPaddress" >> /etc/environment

mkdir /etc/divider-sensors

curl https://raw.githubusercontent.com/byuoitav/divider-sensors-microservice/feature/deployment/config.json > /etc/divider-sensors/
vim /etc/divider-sensors/config.json

curl https://raw.githubusercontent.com/byuoitav/divider-sensors-microservice/feature/deployment/divider-sensors.service > /etc/divider-sensors/
curl https://github.com/byuoitav/divider-sensors-microservice/releases/download/v0.1/divider-sensors-microservice > /etc/divider-sensors/

systemctl enable divider-sensors.service
