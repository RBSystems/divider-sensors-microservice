#!/usr/bin/env bash
read -p 'DSP Microservice Address: ' address
#$ echo 'DSP_MICROSERVICE_ADDRESS='$DSPaddress >> /etc/environment
DSP_MICROSERVICE_ADDRESS="DSP_MICROSERVICE_ADDRESS=$address"
echo $DSP_MICROSERVICE_ADDRESS >> /etc/environment

mkdir /etc/divider-sensors

curl -o /etc/divider-sensors/config.json https://raw.githubusercontent.com/byuoitav/divider-sensors-microservice/feature/deployment/config.json
vi /etc/divider-sensors/config.json

curl -o /etc/systemd/system/divider-sensors.service https://raw.githubusercontent.com/byuoitav/divider-sensors-microservice/feature/deployment/divider-sensors.service
wget  https://github.com/byuoitav/divider-sensors-microservice/releases/download/v0.2/divider-sensors-microservice
chmod 755 divider-sensors-microservice
cp divider-sensors-microservice /etc/divider-sensors

systemctl daemon-reload
systemctl enable divider-sensors.service
systemctl start divider-sensors.service
